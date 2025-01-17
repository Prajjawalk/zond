package transition

import (
	"context"
	"errors"

	"github.com/theQRL/zond/beacon-chain/cache"
	"github.com/theQRL/zond/beacon-chain/state"
	"github.com/theQRL/zond/crypto/hash"
	"github.com/theQRL/zond/encoding/bytesutil"
)

// SkipSlotCache exists for the unlikely scenario that is a large gap between the head state and
// the current slot. If the beacon chain were ever to be stalled for several epochs, it may be
// difficult or impossible to compute the appropriate beacon state for assignments within a
// reasonable amount of time.
var SkipSlotCache = cache.NewSkipSlotCache()

// The key for skip slot cache is mixed between state root and state slot.
// state root is in the mix to defend against different forks with same skip slots
// to hit the same cache. We don't want beacon states mixed up between different chains.
func cacheKey(_ context.Context, state state.ReadOnlyBeaconState) ([32]byte, error) {
	bh := state.LatestBlockHeader()
	if bh == nil {
		return [32]byte{}, errors.New("block head in state can't be nil")
	}
	r, err := bh.HashTreeRoot()
	if err != nil {
		return [32]byte{}, err
	}
	return hash.Hash(append(bytesutil.Bytes32(uint64(state.Slot())), r[:]...)), nil
}
