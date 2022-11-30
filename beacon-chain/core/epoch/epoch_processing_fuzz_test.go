package epoch

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	state_native "github.com/theQRL/zond/beacon-chain/state/state-native"
	ethpb "github.com/theQRL/zond/protos/zond/v1alpha1"
	"github.com/theQRL/zond/testing/require"
)

func TestFuzzFinalUpdates_10000(t *testing.T) {
	fuzzer := fuzz.NewWithSeed(0)
	base := &ethpb.BeaconState{}

	for i := 0; i < 10000; i++ {
		fuzzer.Fuzz(base)
		s, err := state_native.InitializeFromProtoUnsafePhase0(base)
		require.NoError(t, err)
		_, err = ProcessFinalUpdates(s)
		_ = err
	}
}