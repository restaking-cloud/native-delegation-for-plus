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

	if (cfg.K2ContractAddress != common.Address{}) {

		err = e.configureK2Contract(cfg.K2ContractAddress)
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
					"currentGasPrice": currentGasPrice.String()+" gwei",
					"maxGasPrice":     e.cfg.MaxGasPrice.String()+" gwei",
				},
			).Warn("Max gas price is more than 30% lower than current gas price, consider increasing it, else registrations might be paused for a long time")
		}
	}
}

func (e *EthService) FetchProposerRegistryAddressFromK2() (string, error) {

	data, err := e.cfg.K2ContractABI.Pack("proposerRegistry")
	if err != nil {
		return "", err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2ContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	var contractAddress common.Address
	err = e.cfg.K2ContractABI.UnpackIntoInterface(&contractAddress, "proposerRegistry", callResult)
	if err != nil {
		return "", err
	}

	return contractAddress.String(), nil
}

func (e *EthService) BatchCheckRegisteredValidators(validators []phase0.BLSPubKey) (map[string]contracts.BlsPublicKeyToProposerResult, error) {

	var multicallInputs contracts.Multicall3AggregateArgs

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
		fmt.Println("error unpacking batch call result", err)
		return nil, err
	}

	successfulValidatorChecks := make(map[string]contracts.BlsPublicKeyToProposerResult)
	failedValidatorChecks := make(map[string]contracts.BlsPublicKeyToProposerResult)

	for i, validator := range validators {
		if !batchCallResultDecoded.ReturnData[i].Success {
			failedValidatorChecks[validator.String()] = contracts.BlsPublicKeyToProposerResult{}
			fmt.Println("failed validator check", validator, batchCallResultDecoded.ReturnData[i].ReturnData)
			continue
		}

		var registrationResult contracts.BlsPublicKeyToProposerResult
		err = e.cfg.ProposerRegistryContractABI.UnpackIntoInterface(&registrationResult, "blsPublicKeyToProposer", batchCallResultDecoded.ReturnData[i].ReturnData)
		if err != nil {
			fmt.Println("error unpacking blsPublicKeyToProposer result", err)
			return nil, err
		}
		successfulValidatorChecks[validator.String()] = registrationResult
	}

	results := make(map[string]contracts.BlsPublicKeyToProposerResult)
	for k, v := range successfulValidatorChecks {
		results[k] = v
	}
	for k := range failedValidatorChecks {
		results[k] = contracts.BlsPublicKeyToProposerResult{}
	}

	return results, nil
}

func (e *EthService) BatchK2CheckRegisteredValidators(validators []phase0.BLSPubKey) (map[string]string, error) {

	var multicallInputs contracts.Multicall3AggregateArgs

	for _, validator := range validators {

		data, err := e.cfg.K2ContractABI.Pack("blsPublicKeyToNodeOperator", validator[:])
		if err != nil {
			return nil, err
		}

		multicallInputs.Calls = append(multicallInputs.Calls, contracts.Call3{
			Target:       e.cfg.K2ContractAddress,
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
		fmt.Println("error unpacking batch call result", err)
		return nil, err
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

	results := make(map[string]string)
	for k, v := range successfulValidatorChecks {
		results[k] = v
	}
	for k := range failedValidatorChecks {
		results[k] = common.Address{}.String()
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

		sig_v := uint8(reg.ECDSASignature.V)
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
			V: sig_v,
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

		sig_v := uint8(reg.ECDSASignature.V)
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
			V: sig_v,
			R: sig_r32,
			S: sig_s32,
		})
	}

	data, err := e.cfg.K2ContractABI.Pack("batchNodeOperatorDeposit", blsKeys, feeRecipients, blsSignatures, ecdsaSignatures)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.K2ContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending batch node deposit: %w", err)
	}

	return executedTx, nil
}
