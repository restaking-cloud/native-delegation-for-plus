package balanceverifier

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

type EcdsaSignature struct {
	R string `json:"r"`
	S string `json:"s"`
	V uint8 `json:"v"`
}

type ReportEffectiveBalanceResponse struct {
	Responses []ReportEffectiveBalanceResponseItem `json:"responses"`
}

type ReportEffectiveBalanceResponseItem struct {
	Report                      EffectiveBalanceReport `json:"report"`
	DesignatedVerifierSignature EcdsaSignature         `json:"designatedVerifierSignature"`
}

type EffectiveBalanceReport struct {
	BLSPubKey        phase0.BLSPubKey `json:"blsKey"`
	EffectiveBalance uint64           `json:"effectiveBalance"`
}

type Info struct {
	ChainID uint64 `json:"CHAIN_ID,string"`
}

type ReportEffectiveBalancePayload struct {
	BLSPubKeys        []phase0.BLSPubKey `json:"blsKeys"`
	EffectiveBalances []uint64           `json:"effectiveBalances"`
}
