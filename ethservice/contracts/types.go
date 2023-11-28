package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Call struct {
	Target   common.Address `json:"target"`
	CallData []byte         `json:"callData"`
}

type MulticallArgs struct {
	Calls []Call `json:"calls"`
}

type AggregateResult struct {
	BlockNumber *big.Int `json:"blockNumber"`
	ReturnData  []string `json:"returnData"`
}

type Call3 struct {
	Target       common.Address `json:"target"`
	AllowFailure bool           `json:"allowFailure"`
	CallData     []byte         `json:"callData"`
}

type Result struct {
	Success    bool   `json:"success"`
	ReturnData []byte `json:"returnData"`
}

type Call3Value struct {
	Target       common.Address `json:"target"`
	AllowFailure bool           `json:"allowFailure"`
	Value        *big.Int       `json:"value"`
	CallData     []byte         `json:"callData"`
}

type Multicall3AggregateArgs struct {
	Calls []Call3 `json:"calls"`
}

type Multicall3AggregateResult struct {
	ReturnData []Result `json:"returnData"`
}

type Multicall3AggregateValueResult struct {
	ReturnData []Result `json:"returnData"`
}

type Multicall3BlockAndAggregateArgs struct {
	Calls []Call3 `json:"calls"`
}

type Multicall3BlockAndAggregateResult struct {
	BlockNumber *big.Int `json:"blockNumber"`
	BlockHash   string   `json:"blockHash"`
	ReturnData  []Result `json:"returnData"`
}

type BlsPublicKeyToProposerResult struct {
	ActivationBlock *big.Int       `json:"activationBlock"`
	ExitClaimAmount *big.Int       `json:"exitClaimAmount"`
	ExitBlock       *big.Int       `json:"exitBlock"`
	PayoutRecipient common.Address `json:"payoutRecipient"`
	Representative  common.Address `json:"representative"`
	Status          uint8          `json:"status"`
	ReportCount     uint8          `json:"reportCount"`
	OpenClaim       bool           `json:"openClaim"`
}

func (b *BlsPublicKeyToProposerResult) StatusString() string {
	switch b.Status {
	case 0:
		return "UNREGISTERED"
	case 1:
		return "REGISTERED"
	case 2:
		return "ACTIVE"
	case 3:
		return "EXIT_PENDING"
	case 4:
		return "EXITED"
	case 5:
		return "KICKED"
	case 6:
		return "ACTIVE_OUTSIDE_PAYOUT_POOL"
	default:
		return "UNKNOWN"
	}
}
