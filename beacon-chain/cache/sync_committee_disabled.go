//go:build fuzz

package cache

import (
	"github.com/theQRL/zond/beacon-chain/state"
	types "github.com/theQRL/zond/consensus-types/primitives"
)

// FakeSyncCommitteeCache is a fake `SyncCommitteeCache` to satisfy fuzzing.
type FakeSyncCommitteeCache struct {
}

// NewSyncCommittee initializes and returns a new SyncCommitteeCache.
func NewSyncCommittee() *FakeSyncCommitteeCache {
	return &FakeSyncCommitteeCache{}
}

// CurrentEpochIndexPosition -- fake.
func (s *FakeSyncCommitteeCache) CurrentPeriodIndexPosition(root [32]byte, valIdx types.ValidatorIndex) ([]types.CommitteeIndex, error) {
	return nil, nil
}

// NextEpochIndexPosition -- fake.
func (s *FakeSyncCommitteeCache) NextPeriodIndexPosition(root [32]byte, valIdx types.ValidatorIndex) ([]types.CommitteeIndex, error) {
	return nil, nil
}

// UpdatePositionsInCommittee -- fake.
func (s *FakeSyncCommitteeCache) UpdatePositionsInCommittee(syncCommitteeBoundaryRoot [32]byte, state state.BeaconState) error {
	return nil
}