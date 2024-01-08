package ethservice

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	k2common "github.com/restaking-cloud/native-delegation-for-plus/common"

	"github.com/restaking-cloud/native-delegation-for-plus/ethservice/config"
	"github.com/restaking-cloud/native-delegation-for-plus/ethservice/contracts"
)

type EthService struct {
	client *ethclient.Client
	cfg    config.EthServiceConfig
}

func NewEthService() *EthService {
	return &EthService{}
}

func (e *EthService) Configure(cfg config.EthServiceConfig) error {
	e.cfg = cfg

	err := e.connect(cfg.ExecutionNodeUrl)
	if err != nil {
		return err
	}

	err = e.configureMulticallContract(common.HexToAddress(contracts.MULTICALL3_CONTRACT_ADDRESS))
	if err != nil {
		return err
	}

	if (cfg.K2LendingContractAddress != common.Address{}) && (cfg.K2NodeOperatorContractAddress != common.Address{}) {

		err = e.configureK2LendingContract(cfg.K2LendingContractAddress)
		if err != nil {
			return err
		}

		err = e.configureK2NodeOperatorContract(cfg.K2NodeOperatorContractAddress)
		if err != nil {
			return err
		}

		proposerRegistryAddress, err := e.FetchProposerRegistryAddressFromK2()
		if err != nil {
			return err
		}

		// check that the Proposer Registry address matches what is configured if not override
		if proposerRegistryAddress != cfg.ProposerRegistryContractAddress.String() {
			cfg.ProposerRegistryContractAddress = common.HexToAddress(proposerRegistryAddress)
		}

	}

	// check there is a Proposer Registry contract address
	if (cfg.ProposerRegistryContractAddress == common.Address{}) {
		return fmt.Errorf("proposer registry contract address not configured")
		// ideally this would not be possible, since bundled with the module,
		// but in case it happened to be overridden by a wrong address from the K2 contract
	}

	err = e.configureProposerRegistryContract(cfg.ProposerRegistryContractAddress)
	if err != nil {
		return err
	}

	return nil
}

func (e *EthService) ConnectedChainId() *big.Int {
	return e.cfg.ChainID
}

func (e *EthService) Status() (*ethereum.SyncProgress, error) {
	return e.client.SyncProgress(context.Background())
}

func (e *EthService) SetMaxGasPrice(maxGasPrice uint64) {

	e.cfg.MaxGasPrice = big.NewInt(int64(maxGasPrice))

	logger := logrus.WithField("moduleExecution", "k2").WithField("maxGasPrice", e.cfg.MaxGasPrice.String())
	currentGasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.WithError(err).Debug("Failed to retrieve current gas price")
	} else {
		// check if max gas price is more than 30% lower than current gas price
		diff := new(big.Float).Sub(new(big.Float).SetInt(e.cfg.MaxGasPrice), new(big.Float).SetInt(currentGasPrice))
		percentage := new(big.Float).Quo(diff, new(big.Float).SetInt(currentGasPrice))
		if percentage.Cmp(big.NewFloat(-0.3)) < 0 {
			logger.WithFields(
				logrus.Fields{
					"currentGasPrice": currentGasPrice.String() + " gwei",
					"maxGasPrice":     e.cfg.MaxGasPrice.String() + " gwei",
				},
			).Warn("Max gas price is more than 30% lower than current gas price, consider increasing it, else registrations might be paused for a long time")
		}
	}
}

func (e *EthService) FetchProposerRegistryAddressFromK2() (string, error) {

	data, err := e.cfg.K2LendingContractABI.Pack("proposerRegistry")
	if err != nil {
		return "", err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2LendingContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	var contractAddress common.Address
	err = e.cfg.K2LendingContractABI.UnpackIntoInterface(&contractAddress, "proposerRegistry", callResult)
	if err != nil {
		return "", err
	}

	return contractAddress.String(), nil
}

// Proposer Registry Registration

func (e *EthService) BatchCheckRegisteredValidators(validators []phase0.BLSPubKey) (map[string]contracts.BlsPublicKeyToProposerResult, error) {

	var multicallInputs contracts.Multicall3AggregateArgs

	results := make(map[string]contracts.BlsPublicKeyToProposerResult)

	if len(validators) == 0 {
		return results, nil
	}

	for _, validator := range validators {

		data, err := e.cfg.ProposerRegistryContractABI.Pack("blsPublicKeyToProposer", validator[:])
		if err != nil {
			return nil, err
		}

		multicallInputs.Calls = append(multicallInputs.Calls, contracts.Call3{
			Target:       e.cfg.ProposerRegistryContractAddress,
			CallData:     data,
			AllowFailure: true,
		})
	}

	multicallInputsEncoded, err := e.cfg.MulticallContractABI.Pack("aggregate3", multicallInputs.Calls)
	if err != nil {
		return nil, err
	}

	batchCallResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.MulticallContractAddress,
		Data: multicallInputsEncoded,
	}, nil)
	if err != nil {
		return nil, err
	}

	var batchCallResultDecoded contracts.Multicall3AggregateResult
	err = e.cfg.MulticallContractABI.UnpackIntoInterface(&batchCallResultDecoded, "aggregate3", batchCallResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking batch call result: %w", err)
	}

	successfulValidatorChecks := make(map[string]contracts.BlsPublicKeyToProposerResult)
	failedValidatorChecks := make(map[string]contracts.BlsPublicKeyToProposerResult)

	for i, validator := range validators {
		if !batchCallResultDecoded.ReturnData[i].Success {
			failedValidatorChecks[validator.String()] = contracts.BlsPublicKeyToProposerResult{}
			continue
		}

		var registrationResult contracts.BlsPublicKeyToProposerResult
		err = e.cfg.ProposerRegistryContractABI.UnpackIntoInterface(&registrationResult, "blsPublicKeyToProposer", batchCallResultDecoded.ReturnData[i].ReturnData)
		if err != nil {
			return nil, fmt.Errorf("error unpacking blsPublicKeyToProposer result: %w", err)
		}
		successfulValidatorChecks[validator.String()] = registrationResult
	}

	for k, v := range successfulValidatorChecks {
		results[k] = v
	}
	for k := range failedValidatorChecks {
		results[k] = contracts.BlsPublicKeyToProposerResult{}
	}

	return results, nil
}

func (e *EthService) BatchRegisterValidators(validatorRegistrations []k2common.K2ValidatorRegistration) (tx *types.Transaction, err error) {

	var blsKeys [][]byte
	var feeRecipients []common.Address
	var blsSignatures [][]byte
	var ecdsaSignatures []struct {
		V uint8
		R [32]byte
		S [32]byte
	}
	var openClaim []bool

	for _, reg := range validatorRegistrations {

		blsKeys = append(blsKeys, reg.SignedValidatorRegistration.Message.Pubkey[:])
		feeRecipients = append(feeRecipients, common.HexToAddress(reg.SignedValidatorRegistration.Message.FeeRecipient.String()))

		blsSignatures = append(blsSignatures, reg.SignedValidatorRegistration.Signature[:])

		sig_r, err := hex.DecodeString(strings.TrimPrefix(reg.ECDSASignature.R, "0x"))
		if err != nil {
			return nil, err
		}
		var sig_r32 [32]byte
		copy(sig_r32[:], sig_r)
		sig_s, err := hex.DecodeString(strings.TrimPrefix(reg.ECDSASignature.S, "0x"))
		if err != nil {
			return nil, err
		}
		var sig_s32 [32]byte
		copy(sig_s32[:], sig_s)
		ecdsaSignatures = append(ecdsaSignatures, struct {
			V uint8
			R [32]byte
			S [32]byte
		}{
			V: reg.ECDSASignature.V,
			R: sig_r32,
			S: sig_s32,
		})

		openClaim = append(openClaim, true)
	}

	data, err := e.cfg.ProposerRegistryContractABI.Pack("batchRegisterProposerWithoutPayoutPoolRegistration", blsKeys, feeRecipients, blsSignatures, ecdsaSignatures, openClaim, feeRecipients)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.ProposerRegistryContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending batch register: %w", err)
	}

	return executedTx, nil
}

// K2 Native Delegation

func (e *EthService) BatchK2CheckRegisteredValidators(validators []phase0.BLSPubKey) (map[string]string, error) {

	var multicallInputs contracts.Multicall3AggregateArgs

	results := make(map[string]string)

	if len(validators) == 0 {
		return results, nil
	}

	for _, validator := range validators {

		data, err := e.cfg.K2LendingContractABI.Pack("blsPublicKeyToNodeOperator", validator[:])
		if err != nil {
			return nil, err
		}

		multicallInputs.Calls = append(multicallInputs.Calls, contracts.Call3{
			Target:       e.cfg.K2LendingContractAddress,
			CallData:     data,
			AllowFailure: true,
		})
	}

	multicallInputsEncoded, err := e.cfg.MulticallContractABI.Pack("aggregate3", multicallInputs.Calls)
	if err != nil {
		return nil, err
	}

	batchCallResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.MulticallContractAddress,
		Data: multicallInputsEncoded,
	}, nil)
	if err != nil {
		return nil, err
	}

	var batchCallResultDecoded contracts.Multicall3AggregateResult
	err = e.cfg.MulticallContractABI.UnpackIntoInterface(&batchCallResultDecoded, "aggregate3", batchCallResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking batch call result: %w", err)
	}

	successfulValidatorChecks := make(map[string]string)
	failedValidatorChecks := make(map[string]string)

	for i, validator := range validators {
		if !batchCallResultDecoded.ReturnData[i].Success {
			failedValidatorChecks[validator.String()] = ""
			continue
		}

		successfulValidatorChecks[validator.String()] = common.BytesToAddress(batchCallResultDecoded.ReturnData[i].ReturnData).String()
	}

	for k, v := range successfulValidatorChecks {
		results[k] = v
	}
	for k := range failedValidatorChecks {
		results[k] = common.Address{}.String()
	}

	return results, nil
}

func (e *EthService) K2BatchNativeDelegation(validatorRegistrations []k2common.K2ValidatorRegistration) (tx *types.Transaction, err error) {

	// K2 deposit for native delegation
	var blsKeys [][]byte
	var feeRecipients []common.Address
	var blsSignatures [][]byte
	var ecdsaSignatures []struct {
		V uint8
		R [32]byte
		S [32]byte
	}

	for _, reg := range validatorRegistrations {

		blsKeys = append(blsKeys, reg.SignedValidatorRegistration.Message.Pubkey[:])
		feeRecipients = append(feeRecipients, common.HexToAddress(reg.SignedValidatorRegistration.Message.FeeRecipient.String()))

		blsSignatures = append(blsSignatures, reg.SignedValidatorRegistration.Signature[:])

		sig_r, err := hex.DecodeString(strings.TrimPrefix(reg.ECDSASignature.R, "0x"))
		if err != nil {
			return nil, err
		}
		var sig_r32 [32]byte
		copy(sig_r32[:], sig_r)
		sig_s, err := hex.DecodeString(strings.TrimPrefix(reg.ECDSASignature.S, "0x"))
		if err != nil {
			return nil, err
		}
		var sig_s32 [32]byte
		copy(sig_s32[:], sig_s)
		ecdsaSignatures = append(ecdsaSignatures, struct {
			V uint8
			R [32]byte
			S [32]byte
		}{
			V: reg.ECDSASignature.V,
			R: sig_r32,
			S: sig_s32,
		})
	}

	data, err := e.cfg.K2LendingContractABI.Pack("batchNodeOperatorDeposit", blsKeys, feeRecipients, blsSignatures, ecdsaSignatures)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.K2LendingContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending batch node deposit: %w", err)
	}

	return executedTx, nil
}

// Rewards Claiming

func (e *EthService) BatchK2CheckClaimableRewards(validators []phase0.BLSPubKey) (map[string]uint64, error) {

	var multicallInputs contracts.Multicall3AggregateArgs

	results := make(map[string]uint64)

	if len(validators) == 0 {
		return results, nil
	}

	for _, validator := range validators {

		data, err := e.cfg.K2LendingContractABI.Pack("claimableKETHForNodeOperator", validator[:])
		if err != nil {
			return nil, err
		}

		multicallInputs.Calls = append(multicallInputs.Calls, contracts.Call3{
			Target:       e.cfg.K2LendingContractAddress,
			CallData:     data,
			AllowFailure: true,
		})
	}

	multicallInputsEncoded, err := e.cfg.MulticallContractABI.Pack("aggregate3", multicallInputs.Calls)
	if err != nil {
		return nil, err
	}

	batchCallResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.MulticallContractAddress,
		Data: multicallInputsEncoded,
	}, nil)
	if err != nil {
		return nil, err
	}

	var batchCallResultDecoded contracts.Multicall3AggregateResult
	err = e.cfg.MulticallContractABI.UnpackIntoInterface(&batchCallResultDecoded, "aggregate3", batchCallResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking batch call result: %w", err)
	}

	successfulValidatorChecks := make(map[string]uint64)
	failedValidatorChecks := make(map[string]uint64)

	for i, validator := range validators {
		if !batchCallResultDecoded.ReturnData[i].Success {
			failedValidatorChecks[validator.String()] = 0
			continue
		}

		var claimableRewards *big.Int
		err = e.cfg.K2LendingContractABI.UnpackIntoInterface(&claimableRewards, "claimableKETHForNodeOperator", batchCallResultDecoded.ReturnData[i].ReturnData)
		if err != nil {
			return nil, fmt.Errorf("error unpacking claimableKETHForNodeOperator result: %w", err)
		}

		successfulValidatorChecks[validator.String()] = claimableRewards.Uint64()
	}

	for k, v := range successfulValidatorChecks {
		results[k] = v
	}
	for k, v := range failedValidatorChecks {
		results[k] = v
	}

	return results, nil
}

func (e *EthService) BatchK2ClaimRewards(rewardClaims []k2common.K2Claim) (tx *types.Transaction, err error) {

	var blsKeys [][]byte
	var effectiveBalances []*big.Int
	var ecdsaSignatures []struct {
		V uint8
		R [32]byte
		S [32]byte
	}

	for _, claim := range rewardClaims {

		blsKeys = append(blsKeys, claim.ValidatorPubKey[:])

		effectiveBalances = append(effectiveBalances, big.NewInt(0).SetUint64(claim.EffectiveBalance))

		sig_r, err := hex.DecodeString(strings.TrimPrefix(claim.ECDSASignature.R, "0x"))
		if err != nil {
			return nil, err
		}
		var sig_r32 [32]byte
		copy(sig_r32[:], sig_r)
		sig_s, err := hex.DecodeString(strings.TrimPrefix(claim.ECDSASignature.S, "0x"))
		if err != nil {
			return nil, err
		}
		var sig_s32 [32]byte
		copy(sig_s32[:], sig_s)
		ecdsaSignatures = append(ecdsaSignatures, struct {
			V uint8
			R [32]byte
			S [32]byte
		}{
			V: claim.ECDSASignature.V,
			R: sig_r32,
			S: sig_s32,
		})
	}

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("nodeOperatorClaim", blsKeys, effectiveBalances, ecdsaSignatures)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending batch claim: %w", err)
	}

	return executedTx, nil
}

func (e *EthService) K2Exit(validatorExit k2common.K2Exit) (tx *types.Transaction, err error) {

	blsKey := validatorExit.ValidatorPubKey[:]
	effectiveBalance := big.NewInt(0).SetUint64(validatorExit.EffectiveBalance)

	sig_r, err := hex.DecodeString(strings.TrimPrefix(validatorExit.ECDSASignature.R, "0x"))
	if err != nil {
		return nil, err
	}
	var sig_r32 [32]byte
	copy(sig_r32[:], sig_r)
	sig_s, err := hex.DecodeString(strings.TrimPrefix(validatorExit.ECDSASignature.S, "0x"))
	if err != nil {
		return nil, err
	}
	var sig_s32 [32]byte
	copy(sig_s32[:], sig_s)
	ecdsaSignature := struct {
		V uint8
		R [32]byte
		S [32]byte
	}{
		V: validatorExit.ECDSASignature.V,
		R: sig_r32,
		S: sig_s32,
	}

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("nodeOperatorWithdraw", blsKey, effectiveBalance, ecdsaSignature)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending k2 exit: %w", err)
	}

	return executedTx, nil
}

// K2 Capacity, Limits & Node Operator Inclusion list
func (e *EthService) K2CheckInclusionList(nodeOperatorRepresentative common.Address) (bool, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("isPartOfInclusionList", nodeOperatorRepresentative)
	if err != nil {
		return false, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return false, err
	}

	var callResultDecoded bool
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&callResultDecoded, "isPartOfInclusionList", callResult)
	if err != nil {
		return false, fmt.Errorf("error unpacking isPartOfInclusionList result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) K2CheckInclusionListKeysCount(nodeOperatorRepresentative common.Address) (*big.Int, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("totalNumberOfRegisteredKeysForInclusionListMember", nodeOperatorRepresentative)
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&callResultDecoded, "totalNumberOfRegisteredKeysForInclusionListMember", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking totalNumberOfRegisteredKeysForInclusionListMember result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) IndividualMaxNativeDelegation() (*big.Int, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("MAX_NATIVE_DELEGATION_PER_NODE_OPERATOR")
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&callResultDecoded, "MAX_NATIVE_DELEGATION_PER_NODE_OPERATOR", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking MAX_NATIVE_DELEGATION_PER_NODE_OPERATOR result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) GetTotalNativeDelegationCapacityConsumed() (*big.Int, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("totalOpenNativeDelegationCapacityConsumed")
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&callResultDecoded, "totalOpenNativeDelegationCapacityConsumed", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking totalOpenNativeDelegationCapacityConsumed result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) GlobalMaxNativeDelegation() (*big.Int, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("MAX_OPEN_NATIVE_DELEGATION_CAPACITY")
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&callResultDecoded, "MAX_OPEN_NATIVE_DELEGATION_CAPACITY", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking MAX_OPEN_NATIVE_DELEGATION_CAPACITY result: %w", err)
	}

	return callResultDecoded, nil
}
