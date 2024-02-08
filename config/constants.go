package config

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

// A mapping of chain IDs to their respective K2 known configurations
var K2ConfigConstants = map[uint64]K2Config{
	1: {
		K2LendingContractAddress:        common.HexToAddress("0x7D1e9f343a57bD58436b50Ad9935c128a6cF97DB"),
		K2NodeOperatorContractAddress:        common.HexToAddress("0x8eeC404Ef2d4756C972658629400a39359099EF6"),
		ProposerRegistryContractAddress: common.HexToAddress("0xF7F6D8F8b76E94379034d333f4B5FE1694A32D87"),
		SignatureSwapperUrl: &url.URL{
			Scheme: "https",
			Host:   "signature-swapper.ponrelay.com",
		},
		BalanceVerificationUrl: &url.URL{
			Scheme: "https",
			Host:   "verify-effective-balance.restaking.cloud",
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
	17000: {
		K2LendingContractAddress:        common.HexToAddress("0x44cc5A6f4958A04C42C4ceA497ab7A5e5d809f31"),
		K2NodeOperatorContractAddress:        common.HexToAddress("0xc905e03Af4a5A1711b39228C121d2aF2BbAaBdEe"),
		ProposerRegistryContractAddress: common.HexToAddress("0x33a12a1cdc00EE02976fE41509A4A053b9DC5555"),
		SignatureSwapperUrl: &url.URL{
			Scheme: "https",
			Host:   "holesky-signature-swapper.ponrelay.com",
		},
		BalanceVerificationUrl: &url.URL{
			Scheme: "https",
			Host:   "verify-effective-balance-holesky.restaking.cloud",
		},
	},

}
