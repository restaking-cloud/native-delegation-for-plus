package beacon

import (
	"context"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

func (b *BeaconService) Status() (res *SyncStatusData, err error) {
	return b.syncProgress(context.Background())
}

func (b *BeaconService) FinalizedValidatorEffectiveBalance(blsKeys []phase0.BLSPubKey) (res map[phase0.BLSPubKey]uint64, err error) {

	res = make(map[phase0.BLSPubKey]uint64)

	validatorInfo, err := b.getValidatorsFinalizedInfo(context.Background(), blsKeys)
	if err != nil {
		return res, err
	}

	for _, v := range validatorInfo {
		res[v.Validator.Pubkey] = v.Validator.EffectiveBalance
	}

	return res, nil

}
