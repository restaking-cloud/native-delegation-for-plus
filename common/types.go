package common

import (
	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/restaking-cloud/native-delegation-for-plus/signatureswapper"
)

type K2ValidatorRegistration struct {
	ECDSASignature              signatureswapper.EcdsaSignature
	RepresentativeAddress       common.Address
	SignedValidatorRegistration *apiv1.SignedValidatorRegistration
	ProposerRegistrySuccess     bool
	K2Success                   bool
}
