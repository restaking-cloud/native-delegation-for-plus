package config

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

// A mapping of chain IDs to their respective K2 known configurations
var K2ConfigConstants = map[uint64]K2Config{
	1: {
		ProposerRegistryContractAddress: common.HexToAddress("0xF7F6D8F8b76E94379034d333f4B5FE1694A32D87"),
		SignatureSwapperUrl: &url.URL{
			Scheme: "https",
			Host:   "signature-swapper.ponrelay.com",
		},
	},
	5: {
		K2LendingContractAddress:        common.HexToAddress("0xEEc98aBa34AB03EC1533D37F5256651b43E32d05"),
		K2NodeOperatorContractAddress:        common.HexToAddress("0x10b37A1A3e3114fe479B2cf962dB8806c941d2Dc"),
		ProposerRegistryContractAddress: common.HexToAddress("0x1643ec804d944Da97d90c013cBaCD1358Cce1bAF"),
		SignatureSwapperUrl: &url.URL{
			Scheme: "https",
			Host:   "goerli-signature-swapper.ponrelay.com",
		},
		BalanceVerificationUrl: &url.URL{
			Scheme: "https",
			Host:   "verify-effective-balance-goerli.restaking.cloud",
		},
	},
}
