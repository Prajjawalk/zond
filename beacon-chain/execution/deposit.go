package execution

import (
	"context"

	"github.com/pkg/errors"
	"github.com/theQRL/zond/beacon-chain/core/blocks"
	"github.com/theQRL/zond/common"
	"github.com/theQRL/zond/config/params"
	ethpb "github.com/theQRL/zond/protos/zond/v1alpha1"
)

// DepositContractAddress returns the deposit contract address for the given chain.
func DepositContractAddress() (string, error) {
	address := params.BeaconConfig().DepositContractAddress
	if address == "" {
		return "", errors.New("valid deposit contract is required")
	}

	if !common.IsHexAddress(address) {
		return "", errors.New("invalid deposit contract address given: " + address)
	}
	return address, nil
}

func (s *Service) processDeposit(ctx context.Context, eth1Data *ethpb.Eth1Data, deposit *ethpb.Deposit) error {
	var err error
	if err := s.preGenesisState.SetEth1Data(eth1Data); err != nil {
		return err
	}
	beaconState, err := blocks.ProcessPreGenesisDeposits(ctx, s.preGenesisState, []*ethpb.Deposit{deposit})
	if err != nil {
		return errors.Wrap(err, "could not process pre-genesis deposits")
	}
	if beaconState != nil && !beaconState.IsNil() {
		s.preGenesisState = beaconState
	}
	return nil
}
