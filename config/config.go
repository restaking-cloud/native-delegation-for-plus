package config

import (
	"crypto/ecdsa"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

type K2Config struct {
	ValidatorWalletPrivateKey *ecdsa.PrivateKey
	ValidatorWalletAddress    common.Address
	Web3SignerUrl             *url.URL
	SignatureSwapperUrl       *url.URL
	BeaconNodeUrl 		   *url.URL
	ExecutionNodeUrl 	   *url.URL
	K2ContractAddress         common.Address
	ProposerRegistryContractAddress common.Address
	PayoutRecipient common.Address // to override the payout recipient for all validators
}

var K2ConfigDefaults = K2Config{
	ValidatorWalletPrivateKey: nil,
	ValidatorWalletAddress:    common.Address{},
	Web3SignerUrl:             nil,
	SignatureSwapperUrl:       nil,
	BeaconNodeUrl:             nil,
	ExecutionNodeUrl:          nil,
	K2ContractAddress:         common.Address{},
	ProposerRegistryContractAddress: common.Address{},
	PayoutRecipient: common.Address{},
}
