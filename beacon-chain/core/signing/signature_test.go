package signing_test

import (
	"testing"
	"time"

	"github.com/theQRL/zond/beacon-chain/core/signing"
	"github.com/theQRL/zond/config/params"
	"github.com/theQRL/zond/crypto/bls"
	"github.com/theQRL/zond/encoding/bytesutil"
	ethpb "github.com/theQRL/zond/protos/zond/v1alpha1"
	"github.com/theQRL/zond/testing/require"
)

func TestVerifyRegistrationSignature(t *testing.T) {
	sk, err := bls.RandKey()
	require.NoError(t, err)
	reg := &ethpb.ValidatorRegistrationV1{
		FeeRecipient: bytesutil.PadTo([]byte("fee"), 20),
		GasLimit:     123456,
		Timestamp:    uint64(time.Now().Unix()),
		Pubkey:       sk.PublicKey().Marshal(),
	}
	d := params.BeaconConfig().DomainApplicationBuilder
	domain, err := signing.ComputeDomain(d, nil, nil)
	require.NoError(t, err)
	sr, err := signing.ComputeSigningRoot(reg, domain)
	require.NoError(t, err)
	sk.Sign(sr[:]).Marshal()

	sReg := &ethpb.SignedValidatorRegistrationV1{
		Message:   reg,
		Signature: sk.Sign(sr[:]).Marshal(),
	}
	require.NoError(t, signing.VerifyRegistrationSignature(sReg))

	sReg.Signature = []byte("bad")
	require.ErrorIs(t, signing.VerifyRegistrationSignature(sReg), signing.ErrSigFailedToVerify)

	sReg.Message = nil
	require.ErrorIs(t, signing.VerifyRegistrationSignature(sReg), signing.ErrNilRegistration)
}