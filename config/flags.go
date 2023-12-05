package config

import (
	"strings"

	cli "github.com/urfave/cli/v2"
)

var (
	WalletPrivateKeyFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "eth1-private-key",
		Usage:    "The private key of the validator wallet",
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
	MaxGasPriceFlag = &cli.Uint64Flag{
		Name:     ModuleName + "." + "max-gas-price",
		Usage:    "The maximum gas price to use for transactions, in Gwei",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
		Value:    10,
	}
	RegistrationOnlyFlag = &cli.BoolFlag{
		Name:     ModuleName + "." + "registration-only",
		Usage:    "Only register the validators in the proposer registry, do not natively delegate",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
)
