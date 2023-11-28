package signatureswapper

import (
	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type EcdsaSignature struct {
	R string `json:"r"`
	S string `json:"s"`
	V uint64 `json:"v"`
}

type SignatureSwapResponse struct {
	OriginalData *apiv1.SignedValidatorRegistration `json:"originalData"`
	EcdsaSignature EcdsaSignature `json:"ecdsaSignature"`
}

type Info struct {
	ChainID                        uint64 `json:"CHAIN_ID,string"`
	BlsDomain                      string `json:"BLS_DOMAIN"`
	GasLimitProposerRegistryDomain uint64 `json:"GAS_LIMIT_PROPOSER_REGISTRY_DOMAIN,string"`
}

type SignatureSwapPayload struct {
	Signature             phase0.BLSSignature          `json:"signature"`
	Message               *apiv1.ValidatorRegistration `json:"message"`
	RepresentativeAddress common.Address               `json:"representativeAddress"`
}
