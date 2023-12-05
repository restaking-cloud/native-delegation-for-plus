package ethservice

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	ethereum "github.com/ethereum/go-ethereum"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (e *EthService) waitTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	logger := logrus.WithField("moduleExecution", "k2").WithField("tx", tx.Hash().Hex())
	for {
		receipt, err := e.client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, ethereum.NotFound) {
			logger.Trace("Transaction not yet mined")
		} else {
			logger.Trace("Receipt retrieval failed", "err", err)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (e *EthService) transact(context context.Context, tx *types.Transaction, pk *ecdsa.PrivateKey) (signedTx *types.Transaction, err error) {

	walletAddress := crypto.PubkeyToAddress(*pk.Public().(*ecdsa.PublicKey))
	
	gasPrice, err := e.client.SuggestGasPrice(context)
	if err != nil {
		return signedTx, fmt.Errorf("failed to retrieve current gas price: %w", err)
	}
	if e.cfg.MaxGasPrice != nil && gasPrice.Cmp(e.cfg.MaxGasPrice) > 0 {
		return signedTx, fmt.Errorf("gas price (%s) is higher than max gas price (%s)", gasPrice.String(), e.cfg.MaxGasPrice.String())
	}
	gasTip, err := e.client.SuggestGasTipCap(context)
	if err != nil {
		return signedTx, fmt.Errorf("failed to suggest gas tip: %w", err)
	}
	gasLimit, err := e.client.EstimateGas(context, ethereum.CallMsg{
		From: walletAddress,
		To:   tx.To(),
		Data: tx.Data(),
		Value: tx.Value(),
	})
	if err != nil {
		return signedTx, fmt.Errorf("failed to estimate gas: %w", err)
	}
	nonce, err := e.client.PendingNonceAt(context, walletAddress)
	if err != nil {
		return signedTx, fmt.Errorf("failed to get nonce: %w", err)
	}

	fullTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   e.cfg.ChainID,
		Nonce:     nonce,
		GasTipCap: gasTip,
		GasFeeCap: gasPrice,
		Gas:       gasLimit,
		Data:      tx.Data(),
		To:        tx.To(),
	})

	signer := types.LatestSignerForChainID(e.cfg.ChainID)

	signedTx, err = types.SignTx(fullTx, signer, pk)
	if err != nil {
		return signedTx, fmt.Errorf("failed to sign tx: %w", err)
	}

	logger := logrus.WithField("moduleExecution", "k2").WithField("tx", signedTx.Hash().Hex())
	var pending bool

	sendErr := e.client.SendTransaction(context, signedTx)
	if err != nil {
		// check if the transaction was already sent using the hash and if in pending ignore the error
		// this is to handle the case where the transaction was sent but the response was lost or returned a bug error
		_, pending, err = e.client.TransactionByHash(context, signedTx.Hash())
		if err != nil {
			return signedTx, fmt.Errorf("failed to send tx: %w", sendErr)
		}
			
	}
	
	logger.WithField("pending", pending).Info("K2 Module EthService: Transaction sent")

	return signedTx, nil

}

func (e *EthService) transactAndWait(context context.Context, tx *types.Transaction, pk *ecdsa.PrivateKey) (executedTx *types.Transaction, err error) {

	executedTx, err = e.transact(context, tx, pk)
	if err != nil {
		return executedTx, err
	}

	logger := logrus.WithField("moduleExecution", "k2").WithField("tx", executedTx.Hash().Hex())

	logger.Info("K2 Module EthService: Waiting for transaction to be mined")

	receipt, err := e.waitTx(context, executedTx)
	if err != nil {
		return executedTx, fmt.Errorf("failed to wait for tx (%s) to be mined: %w", executedTx.Hash().Hex(), err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return executedTx, fmt.Errorf("tx (%s) failed in execution", executedTx.Hash().Hex())
	}

	logger.Info("K2 Module EthService: Transaction executed successfully")

	return executedTx, nil

}