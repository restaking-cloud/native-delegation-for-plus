package web3signer

import (
	apiv1 "github.com/attestantio/go-builder-client/api/v1"
)

const (
	SignPath = "/api/v1/eth2/sign/"
	ListBLSPubKeysPath = "/api/v1/eth2/publicKeys"
	ReloadSignerKeysPath = "/reload"
	UpCheckPath = "/upcheck"


	VALIDATOR_REGISTRATION_ACTION = "VALIDATOR_REGISTRATION"
)

type ValidatorRegistrationPayload struct {
	Type string `json:"type"`
	SigningRoot string `json:"signing_root,omitempty"`
	ValidatorRegistration *apiv1.ValidatorRegistration `json:"validator_registration"`
}
