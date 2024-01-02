package config

import (
	"crypto/ecdsa"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

type K2Config struct {
	ValidatorWalletPrivateKey       *ecdsa.PrivateKey
	ValidatorWalletAddress          common.Address
	Web3SignerUrl                   *url.URL
	SignatureSwapperUrl             *url.URL
	BeaconNodeUrl                   *url.URL
	ExecutionNodeUrl                *url.URL
	K2LendingContractAddress        common.Address
	K2NodeOperatorContractAddress   common.Address
	ProposerRegistryContractAddress common.Address
	BalanceVerificationUrl          *url.URL       // for effective balance reporting for verifiable signatures to claim rewards
	PayoutRecipient                 common.Address // to override the payout recipient for all validators
	ExclusionListFile               string         // to exclude validators from registration or native delegation
	MaxGasPrice                     uint64
	RegistrationOnly                bool
	ListenAddress                   *url.URL
	ClaimThreshold                  float64 // To only claim rewards if the validator has earned more than this threshold (in KETH)
}

var K2ConfigDefaults = K2Config{
	ValidatorWalletPrivateKey:       nil,
	ValidatorWalletAddress:          common.Address{},
	Web3SignerUrl:                   nil,
	SignatureSwapperUrl:             nil,
	BeaconNodeUrl:                   nil,
	ExecutionNodeUrl:                nil,
	K2LendingContractAddress:        common.Address{},
	K2NodeOperatorContractAddress:   common.Address{},
	ProposerRegistryContractAddress: common.Address{},
	BalanceVerificationUrl:          nil,
	PayoutRecipient:                 common.Address{},
	ExclusionListFile:               "",
	MaxGasPrice:                     0,
	RegistrationOnly:                false,
	ListenAddress:                   &url.URL{Scheme: "http", Host: "localhost:10000"},
	ClaimThreshold:                  0.0,
}
