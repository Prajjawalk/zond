package initialsync

import (
	"context"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sirupsen/logrus"
	mock "github.com/theQRL/zond/beacon-chain/blockchain/testing"
	"github.com/theQRL/zond/beacon-chain/db"
	dbtest "github.com/theQRL/zond/beacon-chain/db/testing"
	"github.com/theQRL/zond/beacon-chain/p2p/peers"
	p2pt "github.com/theQRL/zond/beacon-chain/p2p/testing"
	p2pTypes "github.com/theQRL/zond/beacon-chain/p2p/types"
	beaconsync "github.com/theQRL/zond/beacon-chain/sync"
	"github.com/theQRL/zond/cmd/beacon-chain/flags"
	"github.com/theQRL/zond/config/features"
	"github.com/theQRL/zond/config/params"
	"github.com/theQRL/zond/consensus-types/blocks"
	types "github.com/theQRL/zond/consensus-types/primitives"
	"github.com/theQRL/zond/container/slice"
	"github.com/theQRL/zond/crypto/hash"
	"github.com/theQRL/zond/encoding/bytesutil"
	"github.com/theQRL/zond/p2p/enr"
	ethpb "github.com/theQRL/zond/protos/zond/v1alpha1"
	"github.com/theQRL/zond/testing/assert"
	"github.com/theQRL/zond/testing/require"
	"github.com/theQRL/zond/testing/util"
	prysmTime "github.com/theQRL/zond/time"
	"github.com/theQRL/zond/time/slots"
)

type testCache struct {
	sync.RWMutex
	rootCache       map[types.Slot][32]byte
	parentSlotCache map[types.Slot]types.Slot
}

var cache = &testCache{}

type peerData struct {
	blocks         []types.Slot // slots that peer has blocks
	finalizedEpoch types.Epoch
	headSlot       types.Slot
	failureSlots   []types.Slot // slots at which the peer will return an error
	forkedPeer     bool
}

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)

	resetCfg := features.InitWithReset(&features.Flags{
		EnablePeerScorer: true,
	})
	defer resetCfg()

	resetFlags := flags.Get()
	flags.Init(&flags.GlobalFlags{
		BlockBatchLimit:            64,
		BlockBatchLimitBurstFactor: 10,
	})
	defer func() {
		flags.Init(resetFlags)
	}()

	m.Run()
}

func initializeTestServices(t *testing.T, slots []types.Slot, peers []*peerData) (*mock.ChainService, *p2pt.TestP2P, db.Database) {
	cache.initializeRootCache(slots, t)
	beaconDB := dbtest.SetupDB(t)

	p := p2pt.NewTestP2P(t)
	connectPeers(t, p, peers, p.Peers())
	cache.RLock()
	genesisRoot := cache.rootCache[0]
	cache.RUnlock()

	util.SaveBlock(t, context.Background(), beaconDB, util.NewBeaconBlock())

	st, err := util.NewBeaconState()
	require.NoError(t, err)

	return &mock.ChainService{
		State: st,
		Root:  genesisRoot[:],
		DB:    beaconDB,
		FinalizedCheckPoint: &ethpb.Checkpoint{
			Epoch: 0,
		},
		Genesis:        time.Now(),
		ValidatorsRoot: [32]byte{},
	}, p, beaconDB
}

// makeGenesisTime where now is the current slot.
func makeGenesisTime(currentSlot types.Slot) time.Time {
	return prysmTime.Now().Add(-1 * time.Second * time.Duration(currentSlot) * time.Duration(params.BeaconConfig().SecondsPerSlot))
}

// sanity test on helper function
func TestMakeGenesisTime(t *testing.T) {
	currentSlot := types.Slot(64)
	gt := makeGenesisTime(currentSlot)
	require.Equal(t, currentSlot, slots.Since(gt))
}

// helper function for sequences of block slots
func makeSequence(start, end types.Slot) []types.Slot {
	if end < start {
		panic("cannot make sequence where end is before start")
	}
	seq := make([]types.Slot, 0, end-start+1)
	for i := start; i <= end; i++ {
		seq = append(seq, i)
	}
	return seq
}

func (c *testCache) initializeRootCache(reqSlots []types.Slot, t *testing.T) {
	c.Lock()
	defer c.Unlock()

	c.rootCache = make(map[types.Slot][32]byte)
	c.parentSlotCache = make(map[types.Slot]types.Slot)
	parentSlot := types.Slot(0)

	genesisBlock := util.NewBeaconBlock().Block
	genesisRoot, err := genesisBlock.HashTreeRoot()
	require.NoError(t, err)
	c.rootCache[0] = genesisRoot
	parentRoot := genesisRoot
	for _, slot := range reqSlots {
		currentBlock := util.NewBeaconBlock().Block
		currentBlock.Slot = slot
		currentBlock.ParentRoot = parentRoot[:]
		parentRoot, err = currentBlock.HashTreeRoot()
		require.NoError(t, err)
		c.rootCache[slot] = parentRoot
		c.parentSlotCache[slot] = parentSlot
		parentSlot = slot
	}
}

// sanity test on helper function
func TestMakeSequence(t *testing.T) {
	got := makeSequence(3, 5)
	want := []types.Slot{3, 4, 5}
	require.DeepEqual(t, want, got)
}

// Connect peers with local host. This method sets up peer statuses and the appropriate handlers
// for each test peer.
func connectPeers(t *testing.T, host *p2pt.TestP2P, data []*peerData, peerStatus *peers.Status) {
	for _, d := range data {
		connectPeer(t, host, d, peerStatus)
	}
}

// connectPeer connects a peer to a local host.
func connectPeer(t *testing.T, host *p2pt.TestP2P, datum *peerData, peerStatus *peers.Status) peer.ID {
	const topic = "/eth2/beacon_chain/req/beacon_blocks_by_range/1/ssz_snappy"
	p := p2pt.NewTestP2P(t)
	p.SetStreamHandler(topic, func(stream network.Stream) {
		defer func() {
			assert.NoError(t, stream.Close())
		}()

		req := &ethpb.BeaconBlocksByRangeRequest{}
		assert.NoError(t, p.Encoding().DecodeWithMaxLength(stream, req))

		requestedBlocks := makeSequence(req.StartSlot, req.StartSlot.Add((req.Count-1)*req.Step))

		// Expected failure range
		if len(slice.IntersectionSlot(datum.failureSlots, requestedBlocks)) > 0 {
			_, err := stream.Write([]byte{0x01})
			assert.NoError(t, err)
			msg := p2pTypes.ErrorMessage("bad")
			_, err = p.Encoding().EncodeWithMaxLength(stream, &msg)
			assert.NoError(t, err)
			return
		}

		// Determine the correct subset of blocks to return as dictated by the test scenario.
		ss := slice.IntersectionSlot(datum.blocks, requestedBlocks)

		ret := make([]*ethpb.SignedBeaconBlock, 0)
		for _, slot := range ss {
			if (slot - req.StartSlot).Mod(req.Step) != 0 {
				continue
			}
			cache.RLock()
			parentRoot := cache.rootCache[cache.parentSlotCache[slot]]
			cache.RUnlock()
			blk := util.NewBeaconBlock()
			blk.Block.Slot = slot
			blk.Block.ParentRoot = parentRoot[:]
			// If forked peer, give a different parent root.
			if datum.forkedPeer {
				newRoot := hash.Hash(parentRoot[:])
				blk.Block.ParentRoot = newRoot[:]
			}
			ret = append(ret, blk)
			currRoot, err := blk.Block.HashTreeRoot()
			require.NoError(t, err)
			logrus.Tracef("block with slot %d , signing root %#x and parent root %#x", slot, currRoot, parentRoot)
		}

		if uint64(len(ret)) > req.Count {
			ret = ret[:req.Count]
		}

		mChain := &mock.ChainService{Genesis: time.Now(), ValidatorsRoot: [32]byte{}}
		for i := 0; i < len(ret); i++ {
			wsb, err := blocks.NewSignedBeaconBlock(ret[i])
			require.NoError(t, err)
			assert.NoError(t, beaconsync.WriteBlockChunk(stream, mChain, p.Encoding(), wsb))
		}
	})

	p.Connect(host)

	peerStatus.Add(new(enr.Record), p.PeerID(), nil, network.DirOutbound)
	peerStatus.SetConnectionState(p.PeerID(), peers.PeerConnected)
	peerStatus.SetChainState(p.PeerID(), &ethpb.Status{
		ForkDigest:     params.BeaconConfig().GenesisForkVersion,
		FinalizedRoot:  []byte(fmt.Sprintf("finalized_root %d", datum.finalizedEpoch)),
		FinalizedEpoch: datum.finalizedEpoch,
		HeadRoot:       bytesutil.PadTo([]byte("head_root"), 32),
		HeadSlot:       datum.headSlot,
	})

	return p.PeerID()
}

// extendBlockSequence extends block chain sequentially (creating genesis block, if necessary).
func extendBlockSequence(t *testing.T, inSeq []*ethpb.SignedBeaconBlock, size int) []*ethpb.SignedBeaconBlock {
	// Start from the original sequence.
	outSeq := make([]*ethpb.SignedBeaconBlock, len(inSeq)+size)
	copy(outSeq, inSeq)

	// See if genesis block needs to be created.
	startSlot := len(inSeq)
	if len(inSeq) == 0 {
		outSeq[0] = util.NewBeaconBlock()
		outSeq[0].Block.StateRoot = util.Random32Bytes(t)
		startSlot++
		outSeq = append(outSeq, nil)
	}

	// Extend block chain sequentially.
	for slot := startSlot; slot < len(outSeq); slot++ {
		outSeq[slot] = util.NewBeaconBlock()
		outSeq[slot].Block.Slot = types.Slot(slot)
		parentRoot, err := outSeq[slot-1].Block.HashTreeRoot()
		require.NoError(t, err)
		outSeq[slot].Block.ParentRoot = parentRoot[:]
		// Make sure that blocks having the same slot number, produce different hashes.
		// That way different branches/forks will have different blocks for the same slots.
		outSeq[slot].Block.StateRoot = util.Random32Bytes(t)
	}

	return outSeq
}

// connectPeerHavingBlocks connect host with a peer having provided blocks.
func connectPeerHavingBlocks(
	t *testing.T, host *p2pt.TestP2P, blks []*ethpb.SignedBeaconBlock, finalizedSlot types.Slot,
	peerStatus *peers.Status,
) peer.ID {
	p := p2pt.NewTestP2P(t)

	p.SetStreamHandler("/eth2/beacon_chain/req/beacon_blocks_by_range/1/ssz_snappy", func(stream network.Stream) {
		defer func() {
			_err := stream.Close()
			_ = _err
		}()

		req := &ethpb.BeaconBlocksByRangeRequest{}
		assert.NoError(t, p.Encoding().DecodeWithMaxLength(stream, req))

		for i := req.StartSlot; i < req.StartSlot.Add(req.Count*req.Step); i += types.Slot(req.Step) {
			if uint64(i) >= uint64(len(blks)) {
				break
			}
			chain := &mock.ChainService{Genesis: time.Now(), ValidatorsRoot: [32]byte{}}
			wsb, err := blocks.NewSignedBeaconBlock(blks[i])
			require.NoError(t, err)
			require.NoError(t, beaconsync.WriteBlockChunk(stream, chain, p.Encoding(), wsb))
		}
	})

	p.SetStreamHandler("/eth2/beacon_chain/req/beacon_blocks_by_root/1/ssz_snappy", func(stream network.Stream) {
		defer func() {
			_err := stream.Close()
			_ = _err
		}()

		req := new(p2pTypes.BeaconBlockByRootsReq)
		assert.NoError(t, p.Encoding().DecodeWithMaxLength(stream, req))
		if len(*req) == 0 {
			return
		}
		for _, expectedRoot := range *req {
			for _, blk := range blks {
				if root, err := blk.Block.HashTreeRoot(); err == nil && expectedRoot == root {
					log.Printf("Found blocks_by_root: %#x for slot: %v", root, blk.Block.Slot)
					_, err := stream.Write([]byte{0x00})
					assert.NoError(t, err, "Failed to write to stream")
					_, err = p.Encoding().EncodeWithMaxLength(stream, blk)
					assert.NoError(t, err, "Could not send response back")
				}
			}
		}
	})

	p.Connect(host)

	finalizedEpoch := slots.ToEpoch(finalizedSlot)
	headRoot, err := blks[len(blks)-1].Block.HashTreeRoot()
	require.NoError(t, err)

	peerStatus.Add(new(enr.Record), p.PeerID(), nil, network.DirOutbound)
	peerStatus.SetConnectionState(p.PeerID(), peers.PeerConnected)
	peerStatus.SetChainState(p.PeerID(), &ethpb.Status{
		ForkDigest:     params.BeaconConfig().GenesisForkVersion,
		FinalizedRoot:  []byte(fmt.Sprintf("finalized_root %d", finalizedEpoch)),
		FinalizedEpoch: finalizedEpoch,
		HeadRoot:       headRoot[:],
		HeadSlot:       blks[len(blks)-1].Block.Slot,
	})

	return p.PeerID()
}