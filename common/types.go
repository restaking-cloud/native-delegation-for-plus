package common

import (
	"crypto/ecdsa"
	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/restaking-cloud/native-delegation-for-plus/balanceverifier"
	"github.com/restaking-cloud/native-delegation-for-plus/signatureswapper"
)

type ValidatorWallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

type K2ValidatorRegistration struct {
	ECDSASignature              signatureswapper.EcdsaSignature
	RepresentativeAddress       common.Address
	SignedValidatorRegistration *apiv1.SignedValidatorRegistration
	ProposerRegistrySuccess     bool
	K2Success                   bool
}

type ExcludedValidator struct {
	PublicKey                        phase0.BLSPubKey `json:"publicKey"`
	ExcludedFromProposerRegistration bool             `json:"excludedFromProposerRegistration"`
	ExcludedFromNativeDelegation     bool             `json:"excludedFromNativeDelegation"`
}

type CustomPayoutRepresentative struct {
	RepresentativeAddress common.Address `json:"representativeAddress"`
	FeeRecipientAddress   common.Address `json:"feeRecipientAddress"`
}

type K2Claim struct {
	ValidatorPubKey  phase0.BLSPubKey
	ECDSASignature   balanceverifier.EcdsaSignature
	EffectiveBalance uint64
	ClaimAmount      uint64
	ClaimSuccess     bool
}

type K2Exit struct {
	ValidatorPubKey       phase0.BLSPubKey
	ECDSASignature        balanceverifier.EcdsaSignature
	EffectiveBalance      uint64
	ExitSuccess           bool
	RepresentativeAddress common.Address
}
