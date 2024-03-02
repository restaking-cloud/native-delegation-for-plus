package common

import (
	"crypto/ecdsa"
	"encoding/json"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type EcdsaSignature struct {
	R string `json:"r"`
	S string `json:"s"`
	V uint8  `json:"v"`
}

type ValidatorWallet struct {
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	Address    common.Address    `json:"address"`
}

type K2ValidatorRegistration struct {
	ECDSASignature              EcdsaSignature                     `json:"ecdsaSignature"`
	RepresentativeAddress       common.Address                     `json:"representativeAddress"`
	SignedValidatorRegistration *apiv1.SignedValidatorRegistration `json:"signedValidatorRegistration"`
	ProposerRegistrySuccess     bool                               `json:"proposerRegistrySuccess"`
	K2Success                   bool                               `json:"k2Success"`
}

type ValidatorFilter struct {
	PublicKey            phase0.BLSPubKey `json:"publicKey,omitempty"`
	FeeRecipient         common.Address   `json:"feeRecipientAddress,omitempty"`
	ProposerRegistration bool             `json:"allowProposerRegistration"`
	NativeDelegation     bool             `json:"allowNativeDelegation"`
}

type CustomPayoutRepresentative struct {
	RepresentativeAddress common.Address   `json:"representativeAddress"`
	FeeRecipientAddress   common.Address   `json:"feeRecipientAddress,omitempty"`
	PublicKey             phase0.BLSPubKey `json:"publicKey,omitempty"`
}

type K2Claim struct {
	RepresentativeAddress common.Address `json:"representativeAddress"`
	ClaimAmount           uint64         `json:"claimAmount"`

	// Data used internally to claim rewards
	// Reward claiming requires at least one validate balance report
	// to be submitted to the K2 contract to show use of the K2 protocol
	EffectiveBalanceReportSignature EcdsaSignature   `json:"-"`
	EffectiveBalance                uint64           `json:"-"`
	ValidatorPubKey                 phase0.BLSPubKey `json:"-"`
}

type K2Exit struct {
	ValidatorPubKey       phase0.BLSPubKey `json:"validatorPubKey"`
	ECDSASignature        EcdsaSignature   `json:"ecdsaSignature"`
	EffectiveBalance      uint64           `json:"effectiveBalance"`
	ExitSuccess           bool             `json:"exitSuccess"`
	RepresentativeAddress common.Address   `json:"representativeAddress"`
}

type DelegatedValidator struct {
	ValidatorPubKey                 phase0.BLSPubKey `json:"validatorPubKey"`
	RepresentativeAddress           common.Address   `json:"representativeAddress"`
	EffectiveBalance                uint64           `json:"effectiveBalance"`
	EffectiveBalanceReportSignature EcdsaSignature   `json:"effectiveBalanceReportSignature"`
	IncludeBalance                  bool             `json:"-"`
	IncludeReportSignature          bool             `json:"-"`
}

// create a marshalJSON method for the DelegateValidator struct that ignores the Effective Balance and EffectiveBalanceReportSignature field if it is just an empty signature
func (d DelegatedValidator) MarshalJSON() ([]byte, error) {
	if d.IncludeBalance {
		if d.IncludeReportSignature {
			return json.Marshal(struct {
				ValidatorPubKey                 phase0.BLSPubKey `json:"validatorPubKey"`
				RepresentativeAddress           common.Address   `json:"representativeAddress"`
				EffectiveBalance                uint64           `json:"effectiveBalance"`
				EffectiveBalanceReportSignature EcdsaSignature   `json:"effectiveBalanceReportSignature"`
			}{
				ValidatorPubKey:                 d.ValidatorPubKey,
				RepresentativeAddress:           d.RepresentativeAddress,
				EffectiveBalance:                d.EffectiveBalance,
				EffectiveBalanceReportSignature: d.EffectiveBalanceReportSignature,
			})
		}
		return json.Marshal(struct {
			ValidatorPubKey       phase0.BLSPubKey `json:"validatorPubKey"`
			RepresentativeAddress common.Address   `json:"representativeAddress"`
			EffectiveBalance      uint64           `json:"effectiveBalance"`
		}{
			ValidatorPubKey:       d.ValidatorPubKey,
			RepresentativeAddress: d.RepresentativeAddress,
			EffectiveBalance:      d.EffectiveBalance,
		})
	}
	if d.IncludeReportSignature {
		return json.Marshal(struct {
			ValidatorPubKey                 phase0.BLSPubKey `json:"validatorPubKey"`
			RepresentativeAddress           common.Address   `json:"representativeAddress"`
			EffectiveBalanceReportSignature EcdsaSignature   `json:"effectiveBalanceReportSignature"`
		}{
			ValidatorPubKey:                 d.ValidatorPubKey,
			RepresentativeAddress:           d.RepresentativeAddress,
			EffectiveBalanceReportSignature: d.EffectiveBalanceReportSignature,
		})
	}
	return json.Marshal(struct {
		ValidatorPubKey       phase0.BLSPubKey `json:"validatorPubKey"`
		RepresentativeAddress common.Address   `json:"representativeAddress"`
	}{
		ValidatorPubKey:       d.ValidatorPubKey,
		RepresentativeAddress: d.RepresentativeAddress,
	})
}

type NodeRunnerInfo struct {
	RepresentativeAddress common.Address       `json:"representativeAddress"`
	ClaimableRewards      uint64               `json:"claimableRewards"`
	DelegatedValidators   []DelegatedValidator `json:"delegatedValidators"`
	IncludeBalance        bool                 `json:"-"`
}

func (n *NodeRunnerInfo) MarshalJSON() ([]byte, error) {
	if n.IncludeBalance {
		return json.Marshal(struct {
			RepresentativeAddress common.Address       `json:"representativeAddress"`
			ClaimableRewards      uint64               `json:"claimableRewards"`
			DelegatedValidators   []DelegatedValidator `json:"delegatedValidators"`
		}{
			RepresentativeAddress: n.RepresentativeAddress,
			ClaimableRewards:      n.ClaimableRewards,
			DelegatedValidators:   n.DelegatedValidators,
		})
	}
	return json.Marshal(struct {
		RepresentativeAddress common.Address       `json:"representativeAddress"`
		DelegatedValidators   []DelegatedValidator `json:"delegatedValidators"`
	}{
		RepresentativeAddress: n.RepresentativeAddress,
		DelegatedValidators:   n.DelegatedValidators,
	})
}
