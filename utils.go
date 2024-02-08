package k2

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"strconv"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/fsnotify/fsnotify"

	eth1Common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	k2common "github.com/restaking-cloud/native-delegation-for-plus/common"

	"github.com/pon-network/mev-plus/common"
	"github.com/restaking-cloud/native-delegation-for-plus/config"
)

func (k2 *K2Service) parseConfig(moduleFlags common.ModuleFlags) (err error) {
	for flagName, flagValue := range moduleFlags {
		switch flagName {
		case config.WalletPrivateKeyFlag.Name:
			pk, err := crypto.HexToECDSA(strings.TrimPrefix(flagValue, "0x"))
			if err != nil {
				return fmt.Errorf("-%s: invalid wallet private private key %q", config.WalletPrivateKeyFlag.Name, flagValue)
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
		case config.ExclusionListFlag.Name:
			k2.cfg.ExclusionListFile = flagValue
		case config.MaxGasPriceFlag.Name:
			setMaxGasPrice, err := strconv.ParseUint(flagValue, 10, 64)
			if err != nil {
				return fmt.Errorf("-%s: invalid max gas price %q", config.MaxGasPriceFlag.Name, flagValue)
			}
			if setMaxGasPrice <= 0 {
				return fmt.Errorf("-%s: max gas price must be greater than zero", config.MaxGasPriceFlag.Name)
			}
			k2.cfg.MaxGasPrice = setMaxGasPrice
		case config.RegistrationOnlyFlag.Name:
			k2.cfg.RegistrationOnly, err = strconv.ParseBool(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid registration only flag %q", config.RegistrationOnlyFlag.Name, flagValue)
			}
		case config.ListenAddressFlag.Name:
			k2.cfg.ListenAddress, err = k2common.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.ListenAddressFlag.Name, flagValue)
			}
		case config.ClaimThresholdFlag.Name:
			k2.cfg.ClaimThreshold, err = strconv.ParseFloat(flagValue, 64)
			if err != nil {
				return fmt.Errorf("-%s: invalid claim threshold KETH amount %q", config.ClaimThresholdFlag.Name, flagValue)
			}
			// ensure the claim threshold is positive
			if k2.cfg.ClaimThreshold < 0 {
				return fmt.Errorf("-%s: claim threshold KETH amount must be positive", config.ClaimThresholdFlag.Name)
			}
		case config.K2LendingContractAddressFlag.Name:
			k2.cfg.K2LendingContractAddress = eth1Common.HexToAddress(flagValue)
			if k2.cfg.K2LendingContractAddress == (eth1Common.Address{}) {
				return fmt.Errorf("-%s: invalid address %q", config.K2LendingContractAddressFlag.Name, flagValue)
			}
		case config.K2NodeOperatorContractAddressFlag.Name:
			k2.cfg.K2NodeOperatorContractAddress = eth1Common.HexToAddress(flagValue)
			if k2.cfg.K2NodeOperatorContractAddress == (eth1Common.Address{}) {
				return fmt.Errorf("-%s: invalid address %q", config.K2NodeOperatorContractAddressFlag.Name, flagValue)
			}
		case config.ProposerRegistryContractAddressFlag.Name:
			k2.cfg.ProposerRegistryContractAddress = eth1Common.HexToAddress(flagValue)
			if k2.cfg.ProposerRegistryContractAddress == (eth1Common.Address{}) {
				return fmt.Errorf("-%s: invalid address %q", config.ProposerRegistryContractAddressFlag.Name, flagValue)
			}
		case config.SignatureSwapperUrlFlag.Name:
			k2.cfg.SignatureSwapperUrl, err = k2common.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.SignatureSwapperUrlFlag.Name, flagValue)
			}
		case config.BalanceVerificationUrlFlag.Name:
			k2.cfg.BalanceVerificationUrl, err = k2common.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.BalanceVerificationUrlFlag.Name, flagValue)
			}
		default:
			return fmt.Errorf("unknown flag %q", flagName)
		}

	}

	if len(moduleFlags) > 0 {
		k2.lock.Lock()
		k2.configured = true
		k2.lock.Unlock()
	}

	err = k2.checkConfig()
	if err != nil {
		return err
	}

	return nil
}

func (k2 *K2Service) checkConfig() error {

	k2.lock.Lock()
	defer k2.lock.Unlock()

	if !k2.configured {
		// If not set to run, return
		return nil
	}

	k2.configured = false

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

	// check if exclusion list file is set
	if k2.cfg.ExclusionListFile != "" {
		err := k2.readExclusionList(k2.cfg.ExclusionListFile)
		if err != nil {
			return err
		}
	}

	k2.configured = true

	return nil
}

func (k2 *K2Service) readExclusionList(filePath string) error {

	// Read the exclusion list file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open exclusion list file: %w", err)
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read exclusion list file: %w", err)
	}

	var exclusionList []k2common.ExcludedValidator
	err = json.Unmarshal(fileContent, &exclusionList)
	if err != nil {
		return fmt.Errorf("failed to parse exclusion list file: %w", err)
	}

	preparedExclusionList := make(map[string]k2common.ExcludedValidator)

	// Store the exclusion list
	k2.lock.Lock()
	defer k2.lock.Unlock()
	for _, excludedValidator := range exclusionList {
		// check if validator is already in the exclusion list
		if _, ok := preparedExclusionList[excludedValidator.PublicKey.String()]; ok {
			return fmt.Errorf("duplicate validator %s in exclusion list", excludedValidator.PublicKey.String())
		}
		if excludedValidator.PublicKey == (phase0.BLSPubKey{}) {
			return fmt.Errorf("invalid validator %s in exclusion list", excludedValidator.PublicKey.String())
		}
		if !excludedValidator.ExcludedFromNativeDelegation && !excludedValidator.ExcludedFromProposerRegistration {
			return fmt.Errorf("validator in exclusion list %s must be excluded from either native delegation or proposer registration", excludedValidator.PublicKey.String())
		}
		preparedExclusionList[excludedValidator.PublicKey.String()] = excludedValidator
	}

	k2.exclusionList = preparedExclusionList

	if len(k2.exclusionList) > 0 {
		k2.log.Infof("Exclusion list updated with %d validators", len(k2.exclusionList))
	}

	return nil
}

func (k2 *K2Service) watchExclusionList(filePath string) error {

	// Watch the exclusion list file for changes use the k2.done channel to stop the watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create exclusion list file watcher: %w", err)
	}

	// check if the file exists
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			k2.log.Infof("Exclusion list file %s does not exist, not watching", filePath)
			watcher.Close()
			return fmt.Errorf("exclusion list file does not exist")
		}
		watcher.Close()
		return fmt.Errorf("failed to stat exclusion list file: %w", err)
	}

	// get the parent directory of the file
	fileDir := filepath.Dir(filePath)

	err = watcher.Add(fileDir)
	if err != nil {
		watcher.Close()
		return fmt.Errorf("failed to add exclusion list file to watcher: %w", err)
	}

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					k2.log.Debug("Exclusion list file watcher stopped")
					return
				}
				if (event.Op.Has(fsnotify.Write) || event.Op.Has(fsnotify.Create)) && event.Name == filePath {
					err := k2.readExclusionList(filePath)
					if err != nil {
						k2.log.WithError(err).Warn("Failed to read exclusion list file, not updating exclusion list")
					}
				} else if (event.Op.Has(fsnotify.Remove) || event.Op.Has(fsnotify.Rename)) && event.Name == filePath {
					// check if the file was removed
					if _, err := os.Stat(filePath); !os.IsNotExist(err) {
						// file still exists, not removing exclusion list
						continue
					}
					k2.log.Infof("Exclusion list file %s was renamed/removed, clearing exclusion list", filePath)
					k2.lock.Lock()
					k2.exclusionList = make(map[string]k2common.ExcludedValidator)
					k2.lock.Unlock()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					k2.log.Debug("Exclusion list file watcher stopped")
					return
				}
				k2.log.WithError(err).Error("Failed to watch exclusion list file")
			case <-k2.exit:
				k2.log.Debug("Stopping exclusion list file watcher")
				return
			}
		}
	}()

	return nil
}
