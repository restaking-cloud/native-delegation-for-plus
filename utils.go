package k2

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	eth1Common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	k2common "github.com/restaking-cloud/native-delegation-for-plus/common"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pon-pbs/mev-plus/common"
	"github.com/restaking-cloud/native-delegation-for-plus/config"
)

func (k2 *K2Service) parseConfig(moduleFlags common.ModuleFlags) (err error) {
	for flagName, flagValue := range moduleFlags {
		switch flagName {
		case config.WalletPrivateKeyFlag.Name:
			pk, err := crypto.HexToECDSA(strings.TrimPrefix(flagValue, "0x"))
			if err != nil {
				return fmt.Errorf("-%s: invalid wallet private private key %q", config.WalletPrivateKeyFlag.Name, pk)
			}
			k2.cfg.ValidatorWalletPrivateKey = pk

			publicKey := pk.Public()

			walletAddress := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey))

			k2.cfg.ValidatorWalletAddress = walletAddress

		case config.Web3SignerUrlFlag.Name:
			k2.cfg.Web3SignerUrl, err = k2common.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.Web3SignerUrlFlag.Name, flagValue)
			}
		case config.BeaconNodeUrlFlag.Name:
			k2.cfg.BeaconNodeUrl, err = k2common.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.BeaconNodeUrlFlag.Name, flagValue)
			}
		case config.ExecutionNodeUrlFlag.Name:
			k2.cfg.ExecutionNodeUrl, err = k2common.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.ExecutionNodeUrlFlag.Name, flagValue)
			}
		case config.PayoutRecipientFlag.Name:
			k2.cfg.PayoutRecipient = eth1Common.HexToAddress(flagValue)
			if k2.cfg.PayoutRecipient == (eth1Common.Address{}) {
				return fmt.Errorf("-%s: invalid address %q", config.PayoutRecipientFlag.Name, flagValue)
			}
		default:
			return fmt.Errorf("unknown flag %q", flagName)
		}

	}

	err = k2.checkConfig()
	if err != nil {
		return err
	}

	return nil
}

func (k2 *K2Service) checkConfig() error {

	// if none of the flags are set, return
	if k2.cfg.ValidatorWalletPrivateKey == nil && k2.cfg.Web3SignerUrl == nil && k2.cfg.BeaconNodeUrl == nil && k2.cfg.ExecutionNodeUrl == nil {
		// module not configured to run
		return nil
	}

	// check that the execution node url is set
	if k2.cfg.ExecutionNodeUrl == nil {
		return fmt.Errorf("-%s: execution node url is required", config.ExecutionNodeUrlFlag.Name)
	}

	// check that the beacon node url is set
	if k2.cfg.BeaconNodeUrl == nil {
		return fmt.Errorf("-%s: beacon node url is required", config.BeaconNodeUrlFlag.Name)
	}

	// check that the wallet private key is set
	if k2.cfg.ValidatorWalletPrivateKey == nil {
		return fmt.Errorf("-%s: validator wallet private key is required", config.WalletPrivateKeyFlag.Name)
	}

	// check that the web3 signer url is set
	if k2.cfg.Web3SignerUrl == nil && k2.cfg.PayoutRecipient != (eth1Common.Address{}) {
		return fmt.Errorf("-%s: web3 signer url is required in order to use a custom payout recepient", config.Web3SignerUrlFlag.Name)
	}

	return nil
}

func getListOfBLSKeysFromSignedValidatorRegistration(payload []apiv1.SignedValidatorRegistration) (pubkeys []phase0.BLSPubKey, payloadMap map[string]apiv1.SignedValidatorRegistration) {
	payloadMap = make(map[string]apiv1.SignedValidatorRegistration)
	for _, reg := range payload {
		pubkeys = append(pubkeys, reg.Message.Pubkey)
		payloadMap[reg.Message.Pubkey.String()] = reg
	}
	return pubkeys, payloadMap
}


