package k2

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	k2common "github.com/restaking-cloud/native-delegation-for-plus/common"
	"github.com/sirupsen/logrus"
)

func (k2 *K2Service) processValidatorRegistrations(payload []apiv1.SignedValidatorRegistration) ([]k2common.K2ValidatorRegistration, error) {
	k2.lock.Lock()
	defer k2.lock.Unlock()

	if len(payload) == 0 {
		return nil, nil
	}

	// Default to using the primary representative address for the registrations
	var representative k2common.ValidatorWallet = k2.cfg.ValidatorWallets[0]
	var nodeOperatorTopayoutRecipientMapping map[string]common.Address = make(map[string]common.Address)

	var globalMaxNativeDelegation *big.Int = big.NewInt(0)
	var individualMaxNativeDelegation *big.Int = big.NewInt(0)

	var currentGlobalNativeDelegation *big.Int = big.NewInt(0)
	var currentIndividualNativeDelegation *big.Int = big.NewInt(0)

	var isInInclusionList bool

	var preChecksComplete atomic.Bool
	var preChecksError atomic.Value

	// Start go routines to get the global and current individual native delegation
	// if the currentGlobal is at the max then check the currentIndividual and maxIndividual
	go func() {
		if k2.cfg.K2LendingContractAddress != (common.Address{}) {
			// If the module is configured for K2 operations in addition to Proposer Registry operations
			// use a sync wait group to wait for the results of the global and individual native delegation checks
			var wg sync.WaitGroup
			wg.Add(3)

			go func() {
				defer wg.Done()
				var err error
				globalMaxNativeDelegation, err = k2.eth1.GlobalMaxNativeDelegation()
				if err != nil {
					k2.log.WithError(err).Error("failed to get global max native delegation")
					preChecksError.Store(err)
					return
				}
				k2.log.WithField("globalMaxNativeDelegation", globalMaxNativeDelegation.String()).Debug("global max native delegation")
			}()

			go func() {
				defer wg.Done()
				var err error
				currentGlobalNativeDelegation, err = k2.eth1.GetTotalNativeDelegationCapacityConsumed()
				if err != nil {
					k2.log.WithError(err).Error("failed to get current global native delegation")
					preChecksError.Store(err)
					return
				}
				k2.log.WithField("currentGlobalNativeDelegation", currentGlobalNativeDelegation.String()).Debug("current global native delegation")
			}()

			go func() {
				defer wg.Done()
				var err error
				configuredWalletAddresses := make([]common.Address, len(k2.cfg.ValidatorWallets))
				for i, wallet := range k2.cfg.ValidatorWallets {
					configuredWalletAddresses[i] = wallet.Address
				}
				nodeOperatorTopayoutRecipientMapping, err = k2.eth1.NodeOperatorToPayoutRecipient(configuredWalletAddresses)
				if err != nil {
					k2.log.WithError(err).Error("failed to get configured wallet addresses to payout recipient mapping")
					preChecksError.Store(err)
					return
				}
				k2.log.WithField("payoutRecipientCheck", nodeOperatorTopayoutRecipientMapping).Debug("configured wallet addresses to payout recipient mapping")
			}()

			wg.Wait()

			if preChecksError.Load() != nil {
				// error occurred while getting the global and current individual native delegation
				// so cannot proceed with the registrations
				k2.log.WithError(preChecksError.Load().(error)).Error("failed to get global and current individual native delegation")
				return
			}

			// if k2 operations then the right representative address to use for the registrations must be determined
			payloadFeeRecipient := payload[0].Message.FeeRecipient // can do this safely as batch processing would have grouped by feeRecipient
			var representativeFound bool
			// check if there is a strict representative address for the payload's fee recipient
			if useRepAddress, ok := k2.representativeMapping[payloadFeeRecipient.String()]; ok {
				for _, wallet := range k2.cfg.ValidatorWallets {
					if strings.EqualFold(wallet.Address.String(), useRepAddress.String()) {
						representative = wallet
						representativeFound = true
						break
					}
				}
				if !representativeFound {
					// representative not not configured in the module for the payload's fee recipient
					// throw an error since module cant strictly use this wallet
					k2.log.WithFields(logrus.Fields{
						"payloadFeeRecipient": payloadFeeRecipient.String(),
						"representative":      useRepAddress.String(),
					}).Error("strict representative address for payload's fee recipient not configured")
					preChecksError.Store(fmt.Errorf("strict representative address (%s) for payload's fee recipient (%s) not configured", useRepAddress.String(), payloadFeeRecipient.String()))
					return
				} else {
					// If representative configured, check if the representative has been used in the k2 lending contract
					// before, and if so is the fee recipient the same as the payload's fee recipient
					payoutRecipient := nodeOperatorTopayoutRecipientMapping[representative.Address.String()]
					if !strings.EqualFold(payoutRecipient.String(), payloadFeeRecipient.String()) {
						k2.log.WithFields(logrus.Fields{
							"payloadFeeRecipient":  payloadFeeRecipient.String(),
							"representative":       representative.Address.String(),
							"setK2PayoutRecipient": payoutRecipient.String(),
						}).Error("strict representative address for payload's fee recipient does not match the set payout recipient in the contracts for this representative")
						preChecksError.Store(fmt.Errorf("strict representative address (%s) for payload's fee recipient (%s) does not match the set payout recipient (%s) in the contracts for this representative", representative.Address.String(), payloadFeeRecipient.String(), payoutRecipient.String()))
						return
					}

				}

				k2.log.WithFields(logrus.Fields{
					"payloadFeeRecipient": payloadFeeRecipient.String(),
					"representative":      representative.Address.String(),
				}).Debug("using strict representative address for payload's fee recipient")
			} else { // if no strict representative address for the payload's fee recipient then determine the right representative address to use
				var unusedRepresentative k2common.ValidatorWallet = k2common.ValidatorWallet{
					Address: common.Address{},
				} // holds the next available unused representative address configured
				for _, wallet := range k2.cfg.ValidatorWallets { // Check wallets in order of priority set in the configuration
					payoutRecipient := nodeOperatorTopayoutRecipientMapping[wallet.Address.String()]
					if strings.EqualFold(payoutRecipient.String(), payloadFeeRecipient.String()) {
						representative = wallet // representative found for the set feeRecipient continue further native delegation using this representative as payload fee recipients match
						representativeFound = true
						break
					} else if (payoutRecipient == common.Address{}) && (unusedRepresentative.Address == common.Address{}) {
						// representative not found for the feeRecipient check if the wallet has not been used and is available
						// if a fresh native delegation for these validators with this feeRecipient is needed
						unusedRepresentative = wallet
					}
				}
				if !representativeFound {
					// Representative not found in contract for this payload's payout recipient
					// check if there is an avaialable representative that has not been used
					if unusedRepresentative.Address == (common.Address{}) {
						// no representative available for this payload's payout recipient
						// as all wallets configured have a payout recipient in the k2 lending contracts that do not match the payload's fee recipient thus
						// these validators cannot be registered under this node operator's representative address as they do not pay to the specified payout recipient
						k2.log.WithField("payloadFeeRecipient", payloadFeeRecipient.String()).Error("no representative available for this payload's payout recipient. Add another wallet to your config")
						preChecksError.Store(fmt.Errorf("no configured representative available for this payload's fee recipient [%s]; add another wallet to your config", payloadFeeRecipient.String()))
						return
					}
					// use the next available representative address that has not been used
					representative = unusedRepresentative
				}
			}

			if globalMaxNativeDelegation != nil && currentGlobalNativeDelegation != nil && globalMaxNativeDelegation.Cmp(currentGlobalNativeDelegation) <= 0 {
				// global max native delegation has been reached
				// so need to check individual max native delegation
				k2.log.WithField("globalMaxNativeDelegation", globalMaxNativeDelegation.String()).Debug("global max native delegation has been reached")

				isInInclusionList, err := k2.eth1.K2CheckInclusionList(representative.Address)
				if err != nil {
					k2.log.WithError(err).Error("failed to check if validator's representative address is in inclusion list")
					preChecksError.Store(err)
					return
				}

				if !isInInclusionList {
					// validator representative is not in the inclusion list
					// so cannot natively delegate this validator
					k2.log.WithField("validatorRepresentative", representative.Address.String()).Debug("validator's representative is not in the inclusion list")
					// this is not an error and just means that the validator's representative address is not in the inclusion list to exceed the global max native delegation
					preChecksComplete.Store(true)
					return
				}

				// use a sync wait group to wait for the results of the individual native delegation checks
				var wg sync.WaitGroup
				wg.Add(2)

				go func() {
					defer wg.Done()
					var err error
					individualMaxNativeDelegation, err = k2.eth1.IndividualMaxNativeDelegation()
					if err != nil {
						k2.log.WithError(err).Errorf("failed to get individual max native delegation for representative: %s", representative.Address.String())
						preChecksError.Store(err)
						return
					}
					k2.log.WithField("individualMaxNativeDelegation", individualMaxNativeDelegation.String()).Debug("individual max native delegation")
				}()

				go func() {
					defer wg.Done()
					var err error
					currentIndividualNativeDelegation, err := k2.eth1.K2CheckInclusionListKeysCount(representative.Address)
					if err != nil {
						k2.log.WithError(err).Errorf("failed to get current individual native delegation for representative: %s", representative.Address.String())
						preChecksError.Store(err)
						return
					}
					k2.log.WithField("currentIndividualNativeDelegation", currentIndividualNativeDelegation.String()).Debugf("current individual native delegation for representative: %s", representative.Address.String())
				}()

				wg.Wait()

				if preChecksError.Load() != nil {
					// error occurred while getting the individual native delegation
					// so cannot proceed with the registrations
					k2.log.WithError(preChecksError.Load().(error)).Error("failed to get individual native delegation")
					return
				}

				preChecksComplete.Store(true)
			} else {
				// global max native delegation has not been reached
				// so no need to check individual max native delegation
				k2.log.WithField("globalMaxNativeDelegation", globalMaxNativeDelegation.String()).Debug("global max native delegation has not been reached")
				preChecksComplete.Store(true)
			}
		} else {
			// If the module is configured for Proposer Registry operations only
			// then no need to check the global max native delegation
			preChecksComplete.Store(true)
			k2.log.Debug("module is configured for Proposer Registry operations only, K2 capacity and represenative checks not required")
		}

	}()

	// if no recepient is set then use the default primary representative address to to process this registration batch
	// else if there is already a fee recpient set then check if the module has that wallet configured and use that representative address

	validators, payloadMap := k2common.GetListOfBLSKeysFromSignedValidatorRegistration(payload)

	k2.log.WithField("validators", len(validators)).Info("Checking Proposer Registry registrations")
	proposerRegistryResults, err := k2.eth1.BatchCheckRegisteredValidators(validators)
	if err != nil {
		k2.log.WithError(err).Error("failed to check if validators are already proposerRegistry registered")
		return nil, err
	}

	// for registering in the Proposer Registry
	var registrationsToProcess = make(map[string]apiv1.SignedValidatorRegistration)
	var alreadyRegisteredMap = make(map[string]k2common.K2ValidatorRegistration)

	for validator, registered := range proposerRegistryResults {
		if registered.Status == 0 {
			if excludedValidator, ok := k2.exclusionList[validator]; ok {
				if excludedValidator.ExcludedFromProposerRegistration {
					k2.log.WithField("validatorPubKey", validator).Debug("validator is excluded from Proposer Registry registration")
					continue
				}
			}
			registrationsToProcess[validator] = payloadMap[validator]
		} else {
			alreadyRegisteredMap[validator] = k2common.K2ValidatorRegistration{
				RepresentativeAddress:   registered.Representative,
				ProposerRegistrySuccess: true,
				SignedValidatorRegistration: &apiv1.SignedValidatorRegistration{
					Message: &apiv1.ValidatorRegistration{
						Pubkey:       payloadMap[validator].Message.Pubkey,
						GasLimit:     payloadMap[validator].Message.GasLimit,
						FeeRecipient: bellatrix.ExecutionAddress(common.HexToAddress(registered.PayoutRecipient.String())),
						Timestamp:    payloadMap[validator].Message.Timestamp,
					},
					Signature: payloadMap[validator].Signature,
				},
			}
		}
	}
	proposerRegistryAlreadyRegisteredCount := uint64(len(alreadyRegisteredMap))

	// prepare registrations from the registrationsToProcess for the proposer registry
	processValidators, err := k2.prepareRegistrations(registrationsToProcess, representative.Address)
	if err != nil {
		return nil, err
	}

	var k2AlreadyRegisteredCount uint64
	var k2UnsuppportedCount uint64

	if k2.cfg.K2LendingContractAddress != (common.Address{}) {
		// If the module is configured for K2 operations in addition to Proposer Registry operations

		k2.log.WithField("validators", len(validators)).Info("Checking K2 registrations")
		k2RegistrationResults, err := k2.eth1.BatchK2CheckRegisteredValidators(validators)
		if err != nil {
			k2.log.WithError(err).Error("failed to check if validators are already registered")
			return nil, err
		}

		if len(k2RegistrationResults) > 0 {
			// Wait for the global and individual native delegation checks to complete or for there to be an error
			for !preChecksComplete.Load() && preChecksError.Load() == nil {
				k2.log.Debug("waiting for global and individual native delegation checks to complete")
			}
		}

		// if there is an error in the capacity checks then return the error
		if preChecksError.Load() != nil {
			return nil, preChecksError.Load().(error)
		}

		var k2OnlyRegistrationsToProcess []apiv1.SignedValidatorRegistration

		// prepare K2 only registrations
		for validator, registered := range k2RegistrationResults {
			if strings.EqualFold(registered, common.Address{}.String()) {
				// this is a validator that is not registered in the K2 contract
				if registration, ok := alreadyRegisteredMap[validator]; !ok {
					// this is a validator that is not registered in the Proposer Registry
					// check to see if this is being handled by the registrationToProcess
					if _, ok := registrationsToProcess[validator]; !ok {
						// this should never happen unless the validator key has been excluded from the Proposer Registry registration
						if excludedValidator, ok := k2.exclusionList[validator]; ok {
							if !excludedValidator.ExcludedFromProposerRegistration {
								k2.log.WithField("validatorPubKey", validator).Errorf("validator is not registered in the Proposer Registry and is not being handled by the registrationToProcess")
							}
						} else {
							k2.log.WithField("validatorPubKey", validator).Errorf("validator is not registered in the Proposer Registry and is not being handled by the registrationToProcess")
						}
					}
					continue
				} else {
					// this is a validator that is already registered in the Proposer Registry, but not in the K2 contract
					// so it would NOT be in the registrations mapping from the proposerRegistry registrations to process
					// perform further checks on this validator to add it to the registration mapping if it must be natively delegated

					// check if the validator is excluded from native delegation, before adding it as a registration to process
					if excludedValidator, ok := k2.exclusionList[validator]; ok {
						if excludedValidator.ExcludedFromNativeDelegation {
							k2.log.WithField("validatorPubKey", validator).Debug("validator is excluded from native delegation")
							continue
						}
					}

					// check if the representative address is the same as the one configured
					if registration.RepresentativeAddress != representative.Address {
						// representative address is not the same as the one configured for the module
						// so cannot take action on this validator
						k2.log.WithFields(
							logrus.Fields{
								"validatorPubKey":          validator,
								"representativeAddress":    registration.RepresentativeAddress.String(),
								"configuredRepresentative": representative.Address.String(),
							}).Debugf("validator is already registered in the Proposer Registry, but the representative address is not the same as the one configured for this registration")
						k2UnsuppportedCount++
						continue
					}

					// check if the the fee recipient from the node matches that from the contract for
					// that validator's registration message else the signature will not be valid
					if !strings.EqualFold(registration.SignedValidatorRegistration.Message.FeeRecipient.String(), payloadMap[validator].Message.FeeRecipient.String()) {
						// fee recipient has changed between the registration message and the contract
						// so cannot take action on this validator, as the node's signature will not be valid
						// having signed a payload that has a different fee recipient from what the contract has
						k2.log.WithFields(logrus.Fields{
							"validatorPubKey": validator,
							"registryPayout":  registration.SignedValidatorRegistration.Message.FeeRecipient.String(),
							"nodePayout":      payloadMap[validator].Message.FeeRecipient.String(),
						}).Debugf("validator is already registered in the Proposer Registry, but the fee recipient has changed between the registration message and the contract")
						k2UnsuppportedCount++
						continue
					}

					// representative address is the same as the one configured so can take action on this validator

					// They are in the already registered map which has the proposer registry success set to true,
					// has the signed validator registration payload with the fee recipient the payout recipient registered in the
					// proposer registry contract, and the representative address already set.
					// the only field not set is the ecdsa sig for this already registered validator that would be needed
					// for further k2 native delegation

					// Check if there is capacity for this validator to be natively delegated
					// if the global max native delegation has been reached then check the individual max native delegation
					if globalMaxNativeDelegation != nil && currentGlobalNativeDelegation != nil && globalMaxNativeDelegation.Cmp(currentGlobalNativeDelegation) <= 0 {
						// global max native delegation has been reached
						// so need to check individual max native delegation
						if isInInclusionList {
							if individualMaxNativeDelegation != nil && currentIndividualNativeDelegation != nil && individualMaxNativeDelegation.Cmp(currentIndividualNativeDelegation) <= 0 {
								// individual max native delegation has been reached
								// so cannot natively delegate this validator
								k2.log.WithFields(logrus.Fields{
									"validatorPubKey":   validator,
									"individualMax":     individualMaxNativeDelegation.String(),
									"currentIndividual": currentIndividualNativeDelegation.String(),
								}).Debugf("validator is already registered in the Proposer Registry, but the global and individual max native delegation has been reached")
								k2UnsuppportedCount++
								continue
							}
						} else {
							// validator representative address is not in the inclusion list
							// so cannot natively delegate this validator since the max has been reached
							k2.log.WithField("validatorPubKey", validator).Debug("validator representative address is not in the inclusion list")
							k2UnsuppportedCount++
							continue
						}
					}

					k2OnlyRegistrationsToProcess = append(k2OnlyRegistrationsToProcess, *registration.SignedValidatorRegistration)
					// this would be used to generate the signatures for these validators as they wont have signatures in the processValidators mapping yet

					// Add this validator to the processValidators mapping as it would be used for native delegtion
					processValidators[registration.SignedValidatorRegistration.Message.Pubkey.String()] = registration

					// increase the current global and individual native delegation
					currentGlobalNativeDelegation.Add(currentGlobalNativeDelegation, big.NewInt(1))
					currentIndividualNativeDelegation.Add(currentIndividualNativeDelegation, big.NewInt(1))
				}

			} else {
				// this is a validator that is already registered in the K2 contract meaning it is already registered in the Proposer Registry as well
				// just add to the already registered map with the K2Success flag set to true so that this can be skipped since exists
				// this is just necessary to log and return the results
				if _, ok := alreadyRegisteredMap[validator]; !ok {
					alreadyRegisteredMap[validator] = k2common.K2ValidatorRegistration{
						RepresentativeAddress:   common.HexToAddress(registered),
						ProposerRegistrySuccess: true,
						K2Success:               true,
						SignedValidatorRegistration: &apiv1.SignedValidatorRegistration{
							Message: &apiv1.ValidatorRegistration{
								Pubkey:       payloadMap[validator].Message.Pubkey,
								GasLimit:     payloadMap[validator].Message.GasLimit, // ** this may have changed from the gas limit signed as of the previous time of registration
								FeeRecipient: bellatrix.ExecutionAddress(common.HexToAddress(registered)),
								Timestamp:    payloadMap[validator].Message.Timestamp, // ** this is the timestamp of the payload, but not the time stamp of the actual payload that was signed as of previous time of registration
							},
							Signature: payloadMap[validator].Signature,
						},
					}
					// fields marked ** are not obtainable from the contracts as such would use the current payload fields in returning the result for logging/output
				} else {
					r := alreadyRegisteredMap[validator]
					r.K2Success = true
					alreadyRegisteredMap[validator] = r
				}

				k2AlreadyRegisteredCount++
			}
		}

		// Generate the ECDSA Signatures for the validators who are to perform K2 Only native delegation as these would not have been generated by
		// the previous Proposer Only Registration Preparations. If the validator is not in the alreadyREgistered mapping and is in the registrations to
		// process then it would already have an ECDSA Signature from the signature swapper that it woul use for both registration and native delegation

		// Generate the signatures for the K2-only native delegations that were just added to the mapping
		ecdsaSignatures, err := k2.signatureSwapper.BatchGenerateSignature(k2OnlyRegistrationsToProcess, representative.Address)
		if err != nil {
			k2.log.WithError(err).Error("failed to generate signatures from signature swapper for native delegation")
			return nil, err
		}

		for validatorPubKey, ecdsaSig := range ecdsaSignatures {
			r := processValidators[validatorPubKey.String()]
			r.ECDSASignature = ecdsaSig
			processValidators[validatorPubKey.String()] = r
		}

	}

	var proposerRegistrations []k2common.K2ValidatorRegistration
	var k2Registrations []k2common.K2ValidatorRegistration

	for _, processingDetails := range processValidators {
		if processingDetails.ProposerRegistrySuccess {
			// already registered in the Proposer Registry
			// send the registration to the K2 contract only
			// no need to check if excluded as it wont be in the mapping here if excluded from k2-only native delegation
			if k2.cfg.K2LendingContractAddress != (common.Address{}) {
				k2Registrations = append(k2Registrations, processingDetails)
			}
		} else {
			// not registered in the Proposer Registry
			// send the registration to the Proposer Registry
			// and send the registration to the K2 contract
			// need to check if excluded is it may be in the mapping for only proposer registry registration
			proposerRegistrations = append(proposerRegistrations, processingDetails)
			if k2.cfg.K2LendingContractAddress != (common.Address{}) {
				if excludedValidator, ok := k2.exclusionList[processingDetails.SignedValidatorRegistration.Message.Pubkey.String()]; ok {
					if !excludedValidator.ExcludedFromNativeDelegation {
						k2Registrations = append(k2Registrations, processingDetails)
					}
				} else {
					k2Registrations = append(k2Registrations, processingDetails)
				}
			}
		}
	}

	if len(proposerRegistrations) > 0 {
		k2.log.WithFields(logrus.Fields{
			"proposerRegistrations": len(proposerRegistrations),
			"alreadyRegistered":     proposerRegistryAlreadyRegisteredCount,
		}).Infof("Registering %v validators in the Proposer Registry.", len(proposerRegistrations))
		tx, err := k2.eth1.BatchRegisterValidators(proposerRegistrations)
		if err != nil {
			k2.log.WithError(err).Error("failed to register validators in the Proposer Registry")
			return nil, err
		}
		k2.log.WithFields(logrus.Fields{
			"newRegistrations": len(proposerRegistrations),
			"txHash":           tx.Hash().String(),
		}).Info("Proposer Registry registration transaction completed")
		// update the proposerRegistrySuccess status here as no error was returned from execution
		for _, registration := range proposerRegistrations {
			r := processValidators[registration.SignedValidatorRegistration.Message.Pubkey.String()]
			r.ProposerRegistrySuccess = true
			processValidators[registration.SignedValidatorRegistration.Message.Pubkey.String()] = r
		}
	} else {
		k2.log.WithField("alreadyRegistered", proposerRegistryAlreadyRegisteredCount).Info("No new validators to register in the Proposer Registry")
	}

	if len(k2Registrations) > 0 && k2.cfg.K2LendingContractAddress != (common.Address{}) {
		k2.log.WithFields(logrus.Fields{
			"k2Registrations":   len(k2Registrations),
			"alreadyRegistered": k2AlreadyRegisteredCount,
			"unsupported":       k2UnsuppportedCount,
		}).Infof("Registering %v validators in the K2 contract.", len(k2Registrations))
		tx, err := k2.eth1.K2BatchNativeDelegation(k2Registrations)
		if err != nil {
			k2.log.WithError(err).Error("failed to register validators in the K2 contract")
			return nil, err
		}
		k2.log.WithFields(logrus.Fields{
			"newRegistrations": len(k2Registrations),
			"txHash":           tx.Hash().String(),
		}).Info("K2 registration transaction completed")
		// update the k2Register status here as no error was returned from execution
		for _, registration := range k2Registrations {
			r := processValidators[registration.SignedValidatorRegistration.Message.Pubkey.String()]
			r.K2Success = true
			processValidators[registration.SignedValidatorRegistration.Message.Pubkey.String()] = r
		}
	} else if k2.cfg.K2LendingContractAddress != (common.Address{}) {
		k2.log.WithFields(
			logrus.Fields{
				"alreadyRegistered": k2AlreadyRegisteredCount,
				"unsupported":       k2UnsuppportedCount,
			},
		).Info("No new supported validators to register in the K2 contract")
	}

	var results []k2common.K2ValidatorRegistration
	for _, registration := range processValidators {
		results = append(results, registration)
	}

	if len(processValidators) > 0 {
		k2.log.WithFields(
			logrus.Fields{
				"newRegistrations":                   len(processValidators),
				"newRegistrationsInProposerRegistry": len(proposerRegistrations),
				"newRegistrationsInK2":               len(k2Registrations),
			},
		).Info("Validator registrations successfully processed")
	}

	return results, nil
}

func (k2 *K2Service) processClaim(blsKeys []phase0.BLSPubKey) ([]k2common.K2Claim, error) {
	k2.lock.Lock()
	defer k2.lock.Unlock()

	if k2.cfg.K2LendingContractAddress == (common.Address{}) || k2.cfg.K2NodeOperatorContractAddress == (common.Address{}) {
		// module not configured to run
		return nil, fmt.Errorf("module not configured to run K2 contract operations")
	} else if k2.cfg.BalanceVerificationUrl == nil {
		// module not configured to run
		return nil, fmt.Errorf("module not configured to run balance verification operations for claims")
	}

	k2.log.WithField("validators", len(blsKeys)).Info("Checking K2 claims")
	claimable, err := k2.eth1.BatchK2CheckClaimableRewards(blsKeys)
	if err != nil {
		k2.log.WithError(err).Error("failed to check if validators have claimable rewards")
		return nil, err
	}

	var claimmableValidators []phase0.BLSPubKey
	var claimMap = make(map[string]k2common.K2Claim)
	var claimsToProcess []k2common.K2Claim

	var totalClaimed *big.Float = big.NewFloat(0)

	for _, validator := range blsKeys {
		if claimableAmount, ok := claimable[validator.String()]; ok {
			if claimableAmount > uint64(k2.cfg.ClaimThreshold*math.Pow(10, float64(k2common.KETHDecimals))) {
				claim := k2common.K2Claim{
					ValidatorPubKey: validator,
					ClaimAmount:     claimableAmount,
				}
				claimMap[validator.String()] = claim
				claimmableValidators = append(claimmableValidators, validator)
			} else {
				k2.log.WithFields(logrus.Fields{
					"validator":        validator.String(),
					"claimableRewards": claimableAmount,
				}).Debug("Validator rewards are insufficient to claim")
			}
		}
	}

	if len(claimmableValidators) > 0 {

		// if there are claims then get the effective balance of these validators and
		// report to the balance verifier for signatures for the claims
		effectiveBalances, err := k2.beacon.FinalizedValidatorEffectiveBalance(claimmableValidators)
		if err != nil {
			k2.log.WithError(err).Error("failed to get effective balances for validators")
			return nil, err
		}

		var qualifiedValidators map[phase0.BLSPubKey]uint64 = make(map[phase0.BLSPubKey]uint64)
		for validator, effectiveBalance := range effectiveBalances {
			claim := claimMap[validator.String()]
			claim.EffectiveBalance = effectiveBalance
			claimMap[validator.String()] = claim
			// validators with effective balance less than 32 ETH in a claim attempt
			// would be kicked from the protocol. Best to avoid the software automatically
			// causing such to the user without the user's own direct action
			if effectiveBalance == 32000000000 {
				qualifiedValidators[validator] = effectiveBalance
			} else {
				k2.log.WithFields(logrus.Fields{
					"validator":        validator.String(),
					"effectiveBalance": effectiveBalance,
					"claimableRewards": claimable[validator.String()],
				}).Debugf("Validator has effective balance less than 32 ETH, not processing claim")
			}
		}

		if len(qualifiedValidators) == 0 {
			k2.log.WithField("validators", len(claimmableValidators)).Info("No validators with effective balance of 32 ETH")
			return nil, nil
		}

		verifiedEffectiveBalances, err := k2.balanceverifier.ReportEffectiveBalance(qualifiedValidators)
		if err != nil {
			k2.log.WithError(err).Error("failed to get verified effective balances for validators")
			return nil, err
		}

		for v := range qualifiedValidators {
			if effectiveBalanceSig, ok := verifiedEffectiveBalances[v]; ok {
				updatedClaim := claimMap[v.String()]
				updatedClaim.ECDSASignature = effectiveBalanceSig
				claimsToProcess = append(claimsToProcess, updatedClaim)
				claimMap[v.String()] = updatedClaim
				amountDecimal := big.NewFloat(0).Quo(big.NewFloat(float64(updatedClaim.ClaimAmount)), big.NewFloat(math.Pow(10, float64(k2common.KETHDecimals))))
				totalClaimed.Add(totalClaimed, amountDecimal)
			}
		}

		k2.log.WithFields(logrus.Fields{
			"claims": len(claimsToProcess),
			"amount": totalClaimed.String() + " KETH",
		}).Infof("Processing %v claims through K2 module", len(claimMap))
		tx, err := k2.eth1.BatchK2ClaimRewards(claimsToProcess)
		if err != nil {
			k2.log.WithError(err).Error("failed to claim rewards from the K2 contract")
			return nil, err
		}
		k2.log.WithFields(logrus.Fields{
			"claims": len(claimsToProcess),
			"amount": totalClaimed.String() + " KETH",
			"txHash": tx.Hash().String(),
		}).Info("K2 claim transaction completed")
		// update the claim status here as no error was returned from execution
		for _, claim := range claimsToProcess {
			c := claimMap[claim.ValidatorPubKey.String()]
			c.ClaimSuccess = true
			claimMap[claim.ValidatorPubKey.String()] = c
		}
	} else {
		k2.log.WithField("claims", len(claimMap)).Info("No new claims to process")
		return nil, nil
	}

	var results []k2common.K2Claim
	for _, claim := range claimMap {
		results = append(results, claim)
	}

	k2.log.WithFields(logrus.Fields{
		"claimsRequested:":          len(blsKeys),
		"qualifiedClaimsProcessed:": len(claimsToProcess),
		"amount":                    totalClaimed.String() + " KETH",
	}).Info("K2 claims successfully processed")

	return results, nil
}

func (k2 *K2Service) processExit(blsKey phase0.BLSPubKey) (res k2common.K2Exit, err error) {
	k2.lock.Lock()
	defer k2.lock.Unlock()

	if k2.cfg.K2LendingContractAddress == (common.Address{}) || k2.cfg.K2NodeOperatorContractAddress == (common.Address{}) {
		// module not configured to run
		return res, fmt.Errorf("module not configured to run K2 contract operations")
	} else if k2.cfg.BalanceVerificationUrl == nil {
		// module not configured to run
		return res, fmt.Errorf("module not configured to run balance verification operations for exits")
	}

	k2.log.Info("Checking Validator Representative Address")

	k2RegistrationResults, err := k2.eth1.BatchK2CheckRegisteredValidators([]phase0.BLSPubKey{blsKey})
	if err != nil {
		k2.log.WithError(err).Error("failed to check if validator is already registered")
		return res, fmt.Errorf("failed to check if validator is already registered")
	}

	representativeAddress, ok := k2RegistrationResults[blsKey.String()]
	if !ok {
		k2.log.WithField("validator", blsKey.String()).Error("validator is not registered in the K2 contract")
		return res, fmt.Errorf("validator is not registered in the K2 contract")
	}

	var representative k2common.ValidatorWallet = k2.cfg.ValidatorWallets[0]
	for _, wallet := range k2.cfg.ValidatorWallets {
		if strings.EqualFold(wallet.Address.String(), representativeAddress) {
			representative = wallet
			break
		}
	}

	if !strings.EqualFold(representativeAddress, representative.Address.String()) {
		k2.log.WithFields(logrus.Fields{
			"configuredRepresentative": representative.Address.String(),
			"validatorRepresentative":  representativeAddress,
		}).Error("Validator Representative Address does not match configured representative address")
		return res, fmt.Errorf("validator representative address does not match any configured representative address")
	}

	res.ValidatorPubKey = blsKey
	res.RepresentativeAddress = representative.Address

	// if authorized then get the effective balance of this validator and
	// report to the balance verifier for a signature for the exit (un-delegation)
	effectiveBalances, err := k2.beacon.FinalizedValidatorEffectiveBalance([]phase0.BLSPubKey{blsKey})
	if err != nil {
		k2.log.WithError(err).Error("failed to get effective balance for validator")
		return res, fmt.Errorf("failed to get effective balance for validator")
	}

	res.EffectiveBalance = effectiveBalances[blsKey]

	var report = make(map[phase0.BLSPubKey]uint64)
	report[blsKey] = effectiveBalances[blsKey]

	verifiedEffectiveBalances, err := k2.balanceverifier.ReportEffectiveBalance(report)
	if err != nil {
		k2.log.WithError(err).Error("failed to get verified effective balance for validator")
		return res, fmt.Errorf("failed to get verified effective balance for validator")
	}

	res.ECDSASignature = verifiedEffectiveBalances[blsKey]

	k2.log.WithFields(logrus.Fields{
		"validator": blsKey.String(),
	}).Info("Exiting validator from K2 contract")

	tx, err := k2.eth1.K2Exit(res)
	if err != nil {
		k2.log.WithError(err).Error("failed to exit the validator from the K2 contract")
		return res, fmt.Errorf("failed to exit the validator from the K2 contract: %w", err)
	}
	k2.log.WithFields(logrus.Fields{
		"validator": blsKey.String(),
		"txHash":    tx.Hash().String(),
	}).Info("K2 validator exit transaction completed")
	// update the exit status here as no error was returned from execution
	res.ExitSuccess = true

	k2.log.WithFields(logrus.Fields{
		"validator": blsKey.String(),
	}).Info("K2 validator exit successfully processed")

	return res, nil
}

func (k2 *K2Service) batchProcessClaims(blsKeys []phase0.BLSPubKey) ([]k2common.K2Claim, error) {

	if k2.cfg.K2LendingContractAddress == (common.Address{}) {
		// module not configured to run
		return nil, fmt.Errorf("module not configured to run K2 contract operations")
	} else if k2.cfg.BalanceVerificationUrl == nil {
		// module not configured to run
		return nil, fmt.Errorf("module not configured to run balance verification operations for claims")
	}

	// Split the payload into batches of 90 for the sake of gas efficiency
	var batches [][]phase0.BLSPubKey
	for i := 0; i < len(blsKeys); i += 90 {
		end := i + 90
		if end > len(blsKeys) {
			end = len(blsKeys)
		}
		batches = append(batches, blsKeys[i:end])
	}

	var results []k2common.K2Claim
	for i, batch := range batches {
		k2.log.WithFields(logrus.Fields{
			"currentBatchClaims": len(batch),
			"totalClaims":        len(blsKeys),
		}).Infof("Processing %v claims through K2 module [batch %v/%v]", len(batch), i+1, len(batches))
		batchResults, err := k2.processClaim(batch)
		if err != nil {
			return nil, err
		}
		results = append(results, batchResults...)
	}

	return results, nil
}

func (k2 *K2Service) batchProcessValidatorRegistrations(payload []apiv1.SignedValidatorRegistration) ([]k2common.K2ValidatorRegistration, error) {

	if len(payload) == 0 {
		return nil, nil
	}

	var feeRecipientMapping = make(map[string][]apiv1.SignedValidatorRegistration)
	for _, reg := range payload {
		feeRecipientMapping[reg.Message.FeeRecipient.String()] = append(feeRecipientMapping[reg.Message.FeeRecipient.String()], reg)
	}
	// then group in new sortedPayload by feeRecipient
	var sortedPayload []apiv1.SignedValidatorRegistration
	for feeRecipient, regs := range feeRecipientMapping {
		sortedPayload = append(sortedPayload, regs...)
		k2.log.Debugf("Payload contains registration for fee recipient: %s", feeRecipient)
	}
	payload = sortedPayload

	// Group the payload into batches of 90 for gas efficiency, contract calls, and signature swapping,
	// ensuring batches have the same fee recipient
	var batches [][]apiv1.SignedValidatorRegistration

	var currentRecipient string
	var currentBatch []apiv1.SignedValidatorRegistration
	for i, registration := range payload {
		if i > 0 && (len(currentBatch) == 90 || !strings.EqualFold(registration.Message.FeeRecipient.String(), currentRecipient)) {
			// Append the current batch to batches if it's full or the fee recipient changes
			batches = append(batches, currentBatch)
			currentBatch = nil
		}
		// Update the current fee recipient if it's not set
		if currentRecipient == "" {
			currentRecipient = registration.Message.FeeRecipient.String()
		}
		// Append the current registration to the current batch
		currentBatch = append(currentBatch, registration)
	}
	// Append the remaining batch
	if len(currentBatch) > 0 {
		batches = append(batches, currentBatch)
	}

	var results []k2common.K2ValidatorRegistration
	for i, batch := range batches {
		k2.log.WithFields(logrus.Fields{
			"currentBatchRegistrations": len(batch),
			"totalRegistrations":        len(payload),
		}).Infof("Processing %v validator registrations through K2 module [batch %v/%v]", len(batch), i+1, len(batches))
		batchResults, err := k2.processValidatorRegistrations(batch)
		if err != nil {
			return nil, err
		}
		results = append(results, batchResults...)
	}

	return results, nil
}

func (k2 *K2Service) prepareRegistrations(toProcess map[string]apiv1.SignedValidatorRegistration, representative common.Address) (map[string]k2common.K2ValidatorRegistration, error) {
	// Retrieve the signable pubkeys from the web3signer on runtime since these can change whenever the node runner wishes
	// only if web3signer is configured
	var signablePubKeys map[string]bool = make(map[string]bool)
	var err error
	if k2.cfg.Web3SignerUrl != nil {
		signablePubKeys, err = k2.web3Signer.GetPubkeyList()
		if err != nil {
			k2.log.WithError(err).Error("failed to get signable pubkeys")
			return nil, err
		}
	}

	var registrations map[string]k2common.K2ValidatorRegistration = make(map[string]k2common.K2ValidatorRegistration)

	if len(toProcess) == 0 {
		return registrations, nil
	}

	var signedRegistrations []apiv1.SignedValidatorRegistration

	// prepare Proposer Registry registrations
	for validator, registration := range toProcess {
		var PayoutRecipient bellatrix.ExecutionAddress
		var signedRegistration apiv1.SignedValidatorRegistration

		if k2.cfg.PayoutRecipient != (common.Address{}) && common.HexToAddress(registration.Message.FeeRecipient.String()) != k2.cfg.PayoutRecipient {
			// if a custom the payout recipient is configured, different from the payload
			// check if the validator is in the signable keys list
			// to allow for the signing of a new registration with changed payout recipient
			if _, ok := signablePubKeys[validator]; !ok {
				// validator is not in the signable list so cannot generate a new signature
				// then maintain the original payout recipient
				PayoutRecipient = registration.Message.FeeRecipient
				k2.log.WithFields(logrus.Fields{
					"validatorPubKey": validator,
					"payoutRecipient": registration.Message.FeeRecipient.String(),
				}).Debug("validator is not in the signable list so cannot generate a new signature")
			} else {
				// validator is in the signable list so can generate a new signature
				// then use the custom payout recipient
				PayoutRecipient = bellatrix.ExecutionAddress(k2.cfg.PayoutRecipient)
				k2.log.WithFields(logrus.Fields{
					"validatorPubKey": validator,
					"payoutRecipient": k2.cfg.PayoutRecipient.String(),
				}).Debug("validator is in the signable list so can generate a new signature")
			}
		} else { // if a custom the payout recipient is not configured or is the same as the payload
			PayoutRecipient = registration.Message.FeeRecipient
		}

		if k2.cfg.Web3SignerUrl != nil &&
			common.HexToAddress(registration.Message.FeeRecipient.String()) != k2.cfg.PayoutRecipient &&
			k2.cfg.PayoutRecipient != (common.Address{}) &&
			PayoutRecipient != registration.Message.FeeRecipient {
			// if a custom the payout recipient is configured, different from the payload
			// and a web3 signer is configured, sign the registration with the custom payout recipient

			// this condition would only be hit if the signable key check passed and the payout recipint has been
			// set different from the payload through the preceeding step
			signedRegistration, err = k2.web3Signer.SignRegistration(
				PayoutRecipient,
				registration.Message.GasLimit,
				registration.Message.Pubkey,
				registration.Message.Timestamp,
			)
			if err != nil {
				k2.log.WithField("validatorPubKey", validator).WithError(err).Error("failed to sign registration for validator")
				continue
			}
		} else {
			// if a custom the payout recipient is not configured or is the same as the payload or the web3 signer is not configured
			// use the original signed registration
			signedRegistration = registration
		}

		// Create mapping of qualified registrations to proceed
		registrations[signedRegistration.Message.Pubkey.String()] = k2common.K2ValidatorRegistration{
			RepresentativeAddress:       representative,
			SignedValidatorRegistration: &signedRegistration,
		}
		signedRegistrations = append(signedRegistrations, signedRegistration)

	}

	// Generate signatures for the qualified registration messages
	ecdsaSignatures, err := k2.signatureSwapper.BatchGenerateSignature(signedRegistrations, representative)
	if err != nil {
		k2.log.WithError(err).Error("failed to generate signatures from signature swapper for proposer registration")
		return nil, err
	}

	for validatorPubKey, reg := range registrations {
		reg.ECDSASignature = ecdsaSignatures[reg.SignedValidatorRegistration.Message.Pubkey]
		registrations[validatorPubKey] = reg
	}

	return registrations, nil
}
