// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package zond

import (
	"sync/atomic"
	"time"

	"github.com/theQRL/zond/chain"
	"github.com/theQRL/zond/consensus"
	"github.com/theQRL/zond/consensus_old"
	"github.com/theQRL/zond/core"
	"github.com/theQRL/zond/core/txpool"
	"github.com/theQRL/zond/core/vm"
	"github.com/theQRL/zond/ethdb"
	"github.com/theQRL/zond/internal/zondapi"
	"github.com/theQRL/zond/miner"
	"github.com/theQRL/zond/node"
	"github.com/theQRL/zond/ntp"
	"github.com/theQRL/zond/rpc"
	"github.com/theQRL/zond/zond/downloader"
	"github.com/theQRL/zond/zond/filters"
	"github.com/theQRL/zond/zond/zondconfig"
)

type Zond struct {
	pos          *consensus_old.POS
	blockchainV1 *chain.Chain
	blockchain   *core.BlockChain
	APIBackend   *ZondAPIBackend
	handler      *handler
	miner        *miner.Miner
	txPool       *txpool.TxPool

	// DB interfaces
	chainDb ethdb.Database // Block chain database
	engine  consensus.Engine
}

func (s *Zond) APIs() []rpc.API {
	apis := zondapi.GetAPIs(s.APIBackend)

	// Append any APIs exposed explicitly by the consensus engine
	//apis = append(apis, s.engine.APIs(s.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		//{
		//	Namespace: "zond",
		//	Version:   "1.0",
		//	Service:   NewEthereumAPI(s),
		//}, {
		//	Namespace: "zond",
		//	Version:   "1.0",
		//	Service:   downloader.NewDownloaderAPI(s.handler.downloader, s.eventMux),
		//},
		{
			Namespace: "zond",
			Version:   "1.0",
			Service:   filters.NewFilterAPI(s.APIBackend, false, 5*time.Minute),
		}, // {
		//	Namespace: "admin",
		//	Version:   "1.0",
		//	Service:   NewAdminAPI(s),
		//}, {
		//	Namespace: "debug",
		//	Version:   "1.0",
		//	Service:   NewDebugAPI(s),
		//}, {
		//	Namespace: "net",
		//	Version:   "1.0",
		//	Service:   s.netRPCService,
		//}
	}...)
}

func (s *Zond) BlockChainV1() *chain.Chain {
	return s.blockchainV1
}

func NewV1(stack *node.Node, pos *consensus_old.POS) (*Zond, error) {
	z := &Zond{
		pos:          pos,
		blockchainV1: stack.Blockchain(),
	}
	z.APIBackend = &ZondAPIBackend{stack.Config().ExtRPCEnabled(),
		stack.Config().AllowUnprotectedTxs, z, ntp.GetNTP()}
	stack.RegisterAPIs(z.APIs())
	config := zondconfig.Defaults

	// Override the chain config with provided settings.
	var overrides core.ChainOverrides
	if config.OverrideTerminalTotalDifficulty != nil {
		overrides.OverrideTerminalTotalDifficulty = config.OverrideTerminalTotalDifficulty
	}
	if config.OverrideTerminalTotalDifficultyPassed != nil {
		overrides.OverrideTerminalTotalDifficultyPassed = config.OverrideTerminalTotalDifficultyPassed
	}

	//z.shutdownTracker.MarkStartup()
	return z, nil
}

func New(stack *node.Node) (*Zond, error) {
	z := &Zond{
		// pos:          pos,
		// blockchainV1: stack.Blockchain(),
	}
	z.APIBackend = &ZondAPIBackend{stack.Config().ExtRPCEnabled(),
		stack.Config().AllowUnprotectedTxs, z, ntp.GetNTP()}
	stack.RegisterAPIs(z.APIs())
	config := zondconfig.Defaults
	// Assemble the Ethereum object
	chainDb, err := stack.OpenDatabaseWithFreezer("chaindata", config.DatabaseCache, config.DatabaseHandles, config.DatabaseFreezer, "zond/db/chaindata/", false)
	if err != nil {
		return nil, err
	}

	var (
		vmConfig = vm.Config{
			EnablePreimageRecording: config.EnablePreimageRecording,
		}
		cacheConfig = &core.CacheConfig{
			TrieCleanLimit:      config.TrieCleanCache,
			TrieCleanJournal:    stack.ResolvePath(config.TrieCleanCacheJournal),
			TrieCleanRejournal:  config.TrieCleanCacheRejournal,
			TrieCleanNoPrefetch: config.NoPrefetch,
			TrieDirtyLimit:      config.TrieDirtyCache,
			TrieDirtyDisabled:   config.NoPruning,
			TrieTimeLimit:       config.TrieTimeout,
			SnapshotLimit:       config.SnapshotCache,
			Preimages:           config.Preimages,
		}
	)
	// Override the chain config with provided settings.
	var overrides core.ChainOverrides
	if config.OverrideTerminalTotalDifficulty != nil {
		overrides.OverrideTerminalTotalDifficulty = config.OverrideTerminalTotalDifficulty
	}
	if config.OverrideTerminalTotalDifficultyPassed != nil {
		overrides.OverrideTerminalTotalDifficultyPassed = config.OverrideTerminalTotalDifficultyPassed
	}
	z.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, config.Genesis, &overrides, z.engine, vmConfig, &config.TxLookupLimit)
	if err != nil {
		return nil, err
	}
	z.txPool = txpool.NewTxPool(config.TxPool, z.blockchain.Config(), z.blockchain)
	//z.shutdownTracker.MarkStartup()
	return z, nil
}

func (s *Zond) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Zond) Downloader() *downloader.Downloader { return s.handler.downloader }
func (s *Zond) SyncMode() downloader.SyncMode {
	mode, _ := s.handler.chainSync.modeAndLocalHead()
	return mode
}
func (s *Zond) Synced() bool            { return atomic.LoadUint32(&s.handler.acceptTxs) == 1 }
func (s *Zond) SetSynced()              { atomic.StoreUint32(&s.handler.acceptTxs, 1) }
func (s *Zond) ChainDb() ethdb.Database { return s.chainDb }
func (s *Zond) Miner() *miner.Miner     { return s.miner }
