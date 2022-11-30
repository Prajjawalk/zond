package cache_test

import (
	"context"
	"testing"

	"github.com/theQRL/zond/beacon-chain/cache"
	"github.com/theQRL/zond/beacon-chain/state"
	state_native "github.com/theQRL/zond/beacon-chain/state/state-native"
	ethpb "github.com/theQRL/zond/protos/zond/v1alpha1"
	"github.com/theQRL/zond/testing/assert"
	"github.com/theQRL/zond/testing/require"
)

func TestSkipSlotCache_RoundTrip(t *testing.T) {
	ctx := context.Background()
	c := cache.NewSkipSlotCache()

	r := [32]byte{'a'}
	s, err := c.Get(ctx, r)
	require.NoError(t, err)
	assert.Equal(t, state.BeaconState(nil), s, "Empty cache returned an object")

	require.NoError(t, c.MarkInProgress(r))

	s, err = state_native.InitializeFromProtoPhase0(&ethpb.BeaconState{
		Slot: 10,
	})
	require.NoError(t, err)

	c.Put(ctx, r, s)
	c.MarkNotInProgress(r)

	res, err := c.Get(ctx, r)
	require.NoError(t, err)
	assert.DeepEqual(t, res.ToProto(), s.ToProto(), "Expected equal protos to return from cache")
}