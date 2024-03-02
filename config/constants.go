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
		SubgraphUrl: &url.URL{
			Scheme: "https",
			Host:   "api.thegraph.com",
			Path: "/subgraphs/name/restaking-cloud/k2-protocol",
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
		SubgraphUrl: &url.URL{
			Scheme: "https",
			Host:   "api.thegraph.com",
			Path: "/subgraphs/name/restaking-cloud/k2",
		},
	},
	17000: {
		K2LendingContractAddress:        common.HexToAddress("0x4655512B176243Dd161e61a818899324AE4E9323"),
		K2NodeOperatorContractAddress:        common.HexToAddress("0xe7C28eb37802c4015e65a8c55e182A9d5421Cac3"),
		ProposerRegistryContractAddress: common.HexToAddress("0x33a12a1cdc00EE02976fE41509A4A053b9DC5555"),
		SignatureSwapperUrl: &url.URL{
			Scheme: "https",
			Host:   "holesky-signature-swapper.ponrelay.com",
		},
		BalanceVerificationUrl: &url.URL{
			Scheme: "https",
			Host:   "verify-effective-balance-holesky.restaking.cloud",
		},
		SubgraphUrl: &url.URL{
			Scheme: "https",
			Host:   "api.studio.thegraph.com",
			Path: "/query/45760/k2-holesky/version/latest",
		},
	},

}
