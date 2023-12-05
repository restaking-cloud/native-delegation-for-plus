package k2

import (
	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/ethereum/go-ethereum/common"
	k2common "github.com/restaking-cloud/native-delegation-for-plus/common"
	"github.com/sirupsen/logrus"
)

func (k2 *K2Service) processValidatorRegistrations(payload []apiv1.SignedValidatorRegistration) ([]k2common.K2ValidatorRegistration, error) {
	k2.lock.Lock()
	defer k2.lock.Unlock()

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
			k2.lock.Lock()
			if excludedValidator, ok := k2.exclusionList[validator]; ok {
				if excludedValidator.ExcludedFromProposerRegistration {
					k2.log.WithField("validatorPubKey", validator).Debug("validator is excluded from Proposer Registry registration")
					k2.lock.Unlock()
					continue
				}
			}
			k2.lock.Unlock()
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

	// prepare registrations
	registrations, err := k2.prepareRegistrations(registrationsToProcess)
	if err != nil {
		return nil, err
	}

	var k2AlreadyRegisteredCount uint64
	var k2UnsuppportedCount uint64

	if k2.cfg.K2ContractAddress != (common.Address{}) {
		// If the module is configured for K2 operations in addition to Proposer Registry operations

		k2.log.WithField("validators", len(validators)).Info("Checking K2 registrations")
		k2RegistrationResults, err := k2.eth1.BatchK2CheckRegisteredValidators(validators)
		if err != nil {
			k2.log.WithError(err).Error("failed to check if validators are already registered")
			return nil, err
		}
		// prepare K2 only registrations
		for validator, registered := range k2RegistrationResults {
			if registered == (common.Address{}.String()) {
				// this is a validator that is not registered in the K2 contract
				if registration, ok := alreadyRegisteredMap[validator]; !ok {
					// this is a validator that is not registered in the Proposer Registry
					// check to see if this is being handled by the registrationToProcess
					if _, ok := registrationsToProcess[validator]; !ok {
						// this should never happen unless the validator key has been excluded from the Proposer Registry registration
						k2.lock.Lock()
						if excludedValidator, ok := k2.exclusionList[validator]; ok {
							if !excludedValidator.ExcludedFromProposerRegistration {
								k2.log.WithField("validatorPubKey", validator).Errorf("validator is not registered in the Proposer Registry and is not being handled by the registrationToProcess")
							} 
						} else {
							k2.log.WithField("validatorPubKey", validator).Errorf("validator is not registered in the Proposer Registry and is not being handled by the registrationToProcess")
						}
						k2.lock.Unlock()
					}
					continue
				} else {
					// this is a validator that is already registered in the Proposer Registry, but not in the K2 contract

					// check if the validator is excluded from native delegation
					k2.lock.Lock()
					if excludedValidator, ok := k2.exclusionList[validator]; ok {
						if excludedValidator.ExcludedFromNativeDelegation {
							k2.log.WithField("validatorPubKey", validator).Debug("validator is excluded from native delegation")
							k2.lock.Unlock()
							continue
						}
					}
					k2.lock.Unlock()

					// check if the representative address is the same as the one configured
					if registration.RepresentativeAddress != k2.cfg.ValidatorWalletAddress {
						// representative address is not the same as the one configured for the module
						// so cannot take action on this validator
						k2.log.WithField("validatorPubKey", validator).Debugf("validator is already registered in the Proposer Registry, but the representative address is not the same as the one configured")
						k2UnsuppportedCount++
						continue
					}

					// representative address is the same as the one configured so can take action on this validator
					signature, err := k2.signatureSwapper.GenerateSignature(
						*registration.SignedValidatorRegistration,
						k2.cfg.ValidatorWalletAddress,
					)
					if err != nil {
						k2.log.WithField("validatorPubKey", validator).WithError(err).Error("failed to generate signature")
						continue
					}
					registration.ECDSASignature = signature

					registrations[registration.SignedValidatorRegistration.Message.Pubkey.String()] = registration
				}

			} else {
				// this is a validator that is already registered in the K2 contract meaning it is already registered in the Proposer Registry as well
				if _, ok := alreadyRegisteredMap[validator]; !ok {
					alreadyRegisteredMap[validator] = k2common.K2ValidatorRegistration{
						RepresentativeAddress:   common.HexToAddress(registered),
						ProposerRegistrySuccess: true,
						K2Success:               true,
						SignedValidatorRegistration: &apiv1.SignedValidatorRegistration{
							Message: &apiv1.ValidatorRegistration{
								Pubkey:       payloadMap[validator].Message.Pubkey,
								GasLimit:     payloadMap[validator].Message.GasLimit,
								FeeRecipient: bellatrix.ExecutionAddress(common.HexToAddress(registered)),
								Timestamp:    payloadMap[validator].Message.Timestamp,
							},
							Signature: payloadMap[validator].Signature,
						},
					}
				} else {
					r := alreadyRegisteredMap[validator]
					r.K2Success = true
					alreadyRegisteredMap[validator] = r
				}

				k2AlreadyRegisteredCount++
			}
		}
	}

	var proposerRegistrations []k2common.K2ValidatorRegistration
	var k2Registrations []k2common.K2ValidatorRegistration

	for _, registration := range registrations {
		if registration.ProposerRegistrySuccess {
			// already registered in the Proposer Registry
			// send the registration to the K2 contract only
			k2Registrations = append(k2Registrations, registration)
		} else {
			// not registered in the Proposer Registry
			// send the registration to the Proposer Registry
			// and send the registration to the K2 contract
			proposerRegistrations = append(proposerRegistrations, registration)
			k2Registrations = append(k2Registrations, registration)
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
		// update the proposerRegister status here as no error was returned from execution
		for _, registration := range proposerRegistrations {
			r := registrations[registration.SignedValidatorRegistration.Message.Pubkey.String()]
			r.ProposerRegistrySuccess = true
			registrations[registration.SignedValidatorRegistration.Message.Pubkey.String()] = r
		}
	} else {
		k2.log.WithField("alreadyRegistered", proposerRegistryAlreadyRegisteredCount).Info("No new validators to register in the Proposer Registry")
	}

	if len(k2Registrations) > 0 && k2.cfg.K2ContractAddress != (common.Address{}) {
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
			r := registrations[registration.SignedValidatorRegistration.Message.Pubkey.String()]
			r.K2Success = true
			registrations[registration.SignedValidatorRegistration.Message.Pubkey.String()] = r
		}
	} else if k2.cfg.K2ContractAddress != (common.Address{}) {
		k2.log.WithFields(
			logrus.Fields{
				"alreadyRegistered": k2AlreadyRegisteredCount,
				"unsupported":       k2UnsuppportedCount,
			},
		).Info("No new supported validators to register in the K2 contract")
	}

	var results []k2common.K2ValidatorRegistration
	for _, registration := range registrations {
		results = append(results, registration)
	}

	if len(registrations) > 0 {
		k2.log.WithFields(
			logrus.Fields{
				"newRegistrations":             len(registrations),
				"newRegistrationsInProposerRegistry": len(proposerRegistrations),
				"newRegistrationsInK2":       len(k2Registrations),
			},
		).Info("Validator registrations successfully processed")
	}

	return results, nil
}

func (k2 *K2Service) batchProcessValidatorRegistrations(payload []apiv1.SignedValidatorRegistration) ([]k2common.K2ValidatorRegistration, error) {
	// Split the payload into batches of 1000 for the sake of gas efficiency
	var batches [][]apiv1.SignedValidatorRegistration
	for i := 0; i < len(payload); i += 1000 {
		end := i + 1000
		if end > len(payload) {
			end = len(payload)
		}
		batches = append(batches, payload[i:end])
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

func (k2 *K2Service) prepareRegistrations(toProcess map[string]apiv1.SignedValidatorRegistration) (map[string]k2common.K2ValidatorRegistration, error) {
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

	// prepare Proposer Registry registrations
	for validator, registration := range toProcess {
		var PayoutRecipient bellatrix.ExecutionAddress
		if k2.cfg.PayoutRecipient != (common.Address{}) && common.HexToAddress(registration.Message.FeeRecipient.String()) != k2.cfg.PayoutRecipient {
			// if a custom the payout recipient is configured, different from the payload
			// check if the validator is in the signable keys list
			// to allow for the signing of a new registration with changed payout recipient
			if _, ok := signablePubKeys[validator]; !ok {
				// validator is not in the signable list so cannot generate a new signature
				// then maintain the original payout recipient
				PayoutRecipient = registration.Message.FeeRecipient
			}
		} else {
			PayoutRecipient = registration.Message.FeeRecipient
		}

		var signedRegistration apiv1.SignedValidatorRegistration

		if k2.cfg.Web3SignerUrl != nil && common.HexToAddress(registration.Message.FeeRecipient.String()) != k2.cfg.PayoutRecipient && k2.cfg.PayoutRecipient != (common.Address{}) {
			// if a custom the payout recipient is configured, different from the payload
			// and a web3 signer is configured, sign the registration with the custom payout recipient
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

		signature, err := k2.signatureSwapper.GenerateSignature(
			signedRegistration,
			k2.cfg.ValidatorWalletAddress,
		)
		if err != nil {
			k2.log.WithField("validatorPubKey", validator).WithError(err).Error("failed to generate signature")
			continue
		}

		registrations[signedRegistration.Message.Pubkey.String()] = k2common.K2ValidatorRegistration{
			ECDSASignature:              signature,
			RepresentativeAddress:       k2.cfg.ValidatorWalletAddress,
			SignedValidatorRegistration: &signedRegistration,
		}

	}

	return registrations, nil
}
