package mock

import (
	"context"

	"github.com/theQRL/zond/beacon-chain/state"
	types "github.com/theQRL/zond/consensus-types/primitives"
	eth "github.com/theQRL/zond/protos/zond/v1alpha1"
)

// PoolMock is a fake implementation of PoolManager.
type PoolMock struct {
	Exits []*eth.SignedVoluntaryExit
}

// PendingExits --
func (m *PoolMock) PendingExits(_ state.ReadOnlyBeaconState, _ types.Slot, _ bool) []*eth.SignedVoluntaryExit {
	return m.Exits
}

// InsertVoluntaryExit --
func (m *PoolMock) InsertVoluntaryExit(_ context.Context, _ state.ReadOnlyBeaconState, exit *eth.SignedVoluntaryExit) {
	m.Exits = append(m.Exits, exit)
}

// MarkIncluded --
func (*PoolMock) MarkIncluded(_ *eth.SignedVoluntaryExit) {
	panic("implement me")
}
