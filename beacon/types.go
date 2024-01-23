package beacon

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

type GetSyncStatusResponse struct {
	Data *SyncStatusData `json:"data"`
}

type SyncStatusData struct {
	HeadSlot uint64 `json:"head_slot,string"`
	SyncDistance uint64 `json:"sync_distance,string"`
	IsSyncing bool `json:"is_syncing"`
	IsOptimistic bool `json:"is_optimistic"`
	ElOffline bool `json:"el_offline"`
}

type GetValidatorsResponse struct {
	Data []*ValidatorData `json:"data"`
}

type ValidatorData struct {
	Index uint64 `json:"index,string"`
	Balance uint64 `json:"balance,string"`
	Status string `json:"status"`
	Validator *Validator `json:"validator"`
}

type Validator struct {
	Pubkey phase0.BLSPubKey `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	EffectiveBalance uint64 `json:"effective_balance,string"`
	Slashed bool `json:"slashed"`
	ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch,string"`
	ActivationEpoch uint64 `json:"activation_epoch,string"`
	ExitEpoch uint64 `json:"exit_epoch,string"`
	WithdrawableEpoch uint64 `json:"withdrawable_epoch,string"`
}

type HeadEventData struct {
	Slot  uint64 `json:"slot,string"`
	Block string `json:"block"`
	State string `json:"state"`
}