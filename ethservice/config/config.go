package config

import (
	"crypto/ecdsa"
	"net/url"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type EthServiceConfig struct {
	ExecutionNodeUrl *url.URL
	ChainID *big.Int
	K2ContractAddress               common.Address
	ProposerRegistryContractAddress common.Address

	// ABI
	K2ContractABI *abi.ABI
	ProposerRegistryContractABI *abi.ABI

	// Multicall
	MulticallContractAddress common.Address
	MulticallContractABI *abi.ABI

	ValidatorWalletPrivateKey *ecdsa.PrivateKey
	ValidatorWalletAddress    common.Address
}