package node

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	logTest "github.com/sirupsen/logrus/hooks/test"
	"github.com/theQRL/zond/beacon-chain/blockchain"
	"github.com/theQRL/zond/beacon-chain/builder"
	statefeed "github.com/theQRL/zond/beacon-chain/core/feed/state"
	"github.com/theQRL/zond/beacon-chain/execution"
	mockExecution "github.com/theQRL/zond/beacon-chain/execution/testing"
	"github.com/theQRL/zond/beacon-chain/monitor"
	"github.com/theQRL/zond/cmd"
	"github.com/theQRL/zond/cmd/beacon-chain/flags"
	fieldparams "github.com/theQRL/zond/config/fieldparams"
	"github.com/theQRL/zond/config/params"
	ethpb "github.com/theQRL/zond/proto/prysm/v1alpha1"
	"github.com/theQRL/zond/runtime"
	"github.com/theQRL/zond/runtime/interop"
	"github.com/theQRL/zond/testing/require"
	"github.com/urfave/cli/v2"
)

// Ensure BeaconNode implements interfaces.
var _ statefeed.Notifier = (*BeaconNode)(nil)

// Test that beacon chain node can close.
func TestNodeClose_OK(t *testing.T) {
	hook := logTest.NewGlobal()
	tmp := fmt.Sprintf("%s/datadirtest2", t.TempDir())

	app := cli.App{}
	set := flag.NewFlagSet("test", 0)
	set.Bool("test-skip-pow", true, "skip pow dial")
	set.String("datadir", tmp, "node data directory")
	set.String("p2p-encoding", "ssz", "p2p encoding scheme")
	set.Bool("demo-config", true, "demo configuration")
	set.String("deposit-contract", "0x0000000000000000000000000000000000000000", "deposit contract address")
	set.String("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A", "fee recipient")
	require.NoError(t, set.Set("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A"))
	cmd.ValidatorMonitorIndicesFlag.Value = &cli.IntSlice{}
	cmd.ValidatorMonitorIndicesFlag.Value.SetInt(1)
	ctx := cli.NewContext(&app, set, nil)

	node, err := New(ctx)
	require.NoError(t, err)

	node.Close()

	require.LogsContain(t, hook, "Stopping beacon node")
}

func TestNodeStart_Ok(t *testing.T) {
	hook := logTest.NewGlobal()
	app := cli.App{}
	tmp := fmt.Sprintf("%s/datadirtest2", t.TempDir())
	set := flag.NewFlagSet("test", 0)
	set.String("datadir", tmp, "node data directory")
	set.String("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A", "fee recipient")
	require.NoError(t, set.Set("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A"))

	ctx := cli.NewContext(&app, set, nil)
	node, err := New(ctx, WithBlockchainFlagOptions([]blockchain.Option{}),
		WithBuilderFlagOptions([]builder.Option{}),
		WithExecutionChainOptions([]execution.Option{}))
	require.NoError(t, err)
	node.services = &runtime.ServiceRegistry{}
	go func() {
		node.Start()
	}()
	time.Sleep(3 * time.Second)
	node.Close()
	require.LogsContain(t, hook, "Starting beacon node")

}

func TestNodeStart_Ok_registerDeterministicGenesisService(t *testing.T) {
	numValidators := uint64(1)
	hook := logTest.NewGlobal()
	app := cli.App{}
	tmp := fmt.Sprintf("%s/datadirtest2", t.TempDir())
	set := flag.NewFlagSet("test", 0)
	set.String("datadir", tmp, "node data directory")
	set.Uint64(flags.InteropNumValidatorsFlag.Name, numValidators, "")
	set.String("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A", "fee recipient")
	require.NoError(t, set.Set("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A"))
	genesisState, _, err := interop.GenerateGenesisState(context.Background(), 0, numValidators)
	require.NoError(t, err, "Could not generate genesis beacon state")
	for i := uint64(1); i < 2; i++ {
		someRoot := [32]byte{}
		someKey := [fieldparams.BLSPubkeyLength]byte{}
		copy(someRoot[:], strconv.Itoa(int(i)))
		copy(someKey[:], strconv.Itoa(int(i)))
		genesisState.Validators = append(genesisState.Validators, &ethpb.Validator{
			PublicKey:                  someKey[:],
			WithdrawalCredentials:      someRoot[:],
			EffectiveBalance:           params.BeaconConfig().MaxEffectiveBalance,
			Slashed:                    false,
			ActivationEligibilityEpoch: 1,
			ActivationEpoch:            1,
			ExitEpoch:                  1,
			WithdrawableEpoch:          1,
		})
		genesisState.Balances = append(genesisState.Balances, params.BeaconConfig().MaxEffectiveBalance)
	}
	genesisBytes, err := genesisState.MarshalSSZ()
	require.NoError(t, err)
	require.NoError(t, os.WriteFile("genesis_ssz.json", genesisBytes, 0666))
	set.String(flags.InteropGenesisStateFlag.Name, "genesis_ssz.json", "")
	ctx := cli.NewContext(&app, set, nil)
	node, err := New(ctx, WithBlockchainFlagOptions([]blockchain.Option{}),
		WithBuilderFlagOptions([]builder.Option{}),
		WithExecutionChainOptions([]execution.Option{}))
	require.NoError(t, err)
	node.services = &runtime.ServiceRegistry{}
	go func() {
		node.Start()
	}()
	time.Sleep(3 * time.Second)
	node.Close()
	require.LogsContain(t, hook, "Starting beacon node")
	require.NoError(t, os.Remove("genesis_ssz.json"))
}

// TestClearDB tests clearing the database
func TestClearDB(t *testing.T) {
	hook := logTest.NewGlobal()
	srv, endpoint, err := mockExecution.SetupRPCServer()
	require.NoError(t, err)
	t.Cleanup(func() {
		srv.Stop()
	})

	tmp := filepath.Join(t.TempDir(), "datadirtest")

	app := cli.App{}
	set := flag.NewFlagSet("test", 0)
	set.String("datadir", tmp, "node data directory")
	set.Bool(cmd.ForceClearDB.Name, true, "force clear db")
	set.String("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A", "fee recipient")
	require.NoError(t, set.Set("suggested-fee-recipient", "0x6e35733c5af9B61374A128e6F85f553aF09ff89A"))
	context := cli.NewContext(&app, set, nil)
	_, err = New(context, WithExecutionChainOptions([]execution.Option{
		execution.WithHttpEndpoint(endpoint),
	}))
	require.NoError(t, err)

	require.LogsContain(t, hook, "Removing database")
}

func TestMonitor_RegisteredCorrectly(t *testing.T) {
	app := cli.App{}
	set := flag.NewFlagSet("test", 0)
	require.NoError(t, cmd.ValidatorMonitorIndicesFlag.Apply(set))
	cliCtx := cli.NewContext(&app, set, nil)
	require.NoError(t, cliCtx.Set(cmd.ValidatorMonitorIndicesFlag.Name, "1,2"))
	n := &BeaconNode{ctx: context.Background(), cliCtx: cliCtx, services: runtime.NewServiceRegistry()}
	require.NoError(t, n.services.RegisterService(&blockchain.Service{}))
	require.NoError(t, n.registerValidatorMonitorService())

	var mService *monitor.Service
	require.NoError(t, n.services.FetchService(&mService))
	require.Equal(t, true, mService.TrackedValidators[1])
	require.Equal(t, true, mService.TrackedValidators[2])
	require.Equal(t, false, mService.TrackedValidators[100])
}
