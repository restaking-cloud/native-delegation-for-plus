package config

import (
	"strings"

	cli "github.com/urfave/cli/v2"
)

const ModuleName = "k2"

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:      ModuleName,
		Usage:     "Start the K2 native delegation module",
		UsageText: "The K2 native delegation module is responsible for on-chain registration of native delegations and proposer registries",
		Category:  strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
		Flags:     k2Flags(),
	}
}

func k2Flags() []cli.Flag {
	return []cli.Flag{
		WalletPrivateKeyFlag,
		Web3SignerUrlFlag,
		PayoutRecipientFlag,
		BeaconNodeUrlFlag,
		ExecutionNodeUrlFlag,
		ExclusionListFlag,
		MaxGasPriceFlag,
		RegistrationOnlyFlag,
		ListenAddressFlag,
		ClaimThresholdFlag,
	}
}