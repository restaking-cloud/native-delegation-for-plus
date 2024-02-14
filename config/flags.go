package config

import (
	"strings"

	cli "github.com/urfave/cli/v2"
)

var (
	WalletPrivateKeyFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "eth1-private-key",
		Usage:    "The private key of the validator wallet. You can load multiple keys by separating them with a comma in order of priority",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
		EnvVars:  []string{"ETH1_PRIVATE_KEY"},
	}
	Web3SignerUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "web3-signer-url",
		Usage:    "The url of the web3 signer",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	PayoutRecipientFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "payout-recipient",
		Usage:    "The address of the payout recipient, optional would then default to the registration fee recepient address",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	BeaconNodeUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "beacon-node-url",
		Usage:    "The url of the beacon node",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	ExecutionNodeUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "execution-node-url",
		Usage:    "The url of the execution node",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	ExclusionListFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "exclusion-list",
		Usage:    "The list of addresses to exclude from either the Proposer Registration or Native Delegation",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	RepresentativeMappingFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "representative-mapping",
		Usage:    "The mapping of representative addresses designated to handle validators that pay to different fee recipients",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	MaxGasPriceFlag = &cli.Uint64Flag{
		Name:     ModuleName + "." + "max-gas-price",
		Usage:    "The maximum gas price to use for transactions, in Wei",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	RegistrationOnlyFlag = &cli.BoolFlag{
		Name:     ModuleName + "." + "registration-only",
		Usage:    "Only register the validators in the proposer registry, do not natively delegate",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	ListenAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "listen-address",
		Usage:    "The address to listen on for incoming requests",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	ClaimThresholdFlag = &cli.Float64Flag{
		Name:     ModuleName + "." + "claim-threshold",
		Usage:    "The threshold for claiming rewards, in KETH",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	K2LendingContractAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "k2-lending-contract-address",
		Usage:    "The address of the K2 lending contract to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	K2NodeOperatorContractAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "k2-node-operator-contract-address",
		Usage:    "The address of the K2 node operator contract to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	ProposerRegistryContractAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "proposer-registry-contract-address",
		Usage:    "The address of the proposer registry contract to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	SignatureSwapperUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "signature-swapper-url",
		Usage:    "The url of the signature swapper to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	BalanceVerificationUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "balance-verification-url",
		Usage:    "The url of the balance verification service to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
)
