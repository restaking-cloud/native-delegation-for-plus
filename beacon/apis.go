package beacon

import (
	"fmt"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/r3labs/sse/v2"
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

func (b *BeaconService) SubscribeToHeadEvents(ctx context.Context, headEvent chan<- HeadEventData) error {
	/*
		Subscribe to head events from the beacon chain
		Events are sent to the headChannel
	*/
	logger := logrus.WithField("moduleExecution", "k2")
	logger.Debugf("Starting head events subscription to node:%s", b.cfg.BeaconNodeUrl.String())
	defer logger.Debugf("Head events subscription ended")

	for {
		client := sse.NewClient(fmt.Sprintf("%s/eth/v1/events?topics=head", b.cfg.BeaconNodeUrl.String()))
		// Use sse client to subscribe to events
		err := client.SubscribeRawWithContext(ctx, func(msg *sse.Event) {
			var event HeadEventData
			if err := json.Unmarshal(msg.Data, &event); err != nil {
				logger.Warn("Head event subscription failed", "error", err)
				return
			}
			b.mu.Lock()
			if b.currentSlot < event.Slot {
				b.currentSlot = event.Slot
			}
			b.mu.Unlock()
			headEvent <- event
		})

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return nil
		}

		if err != nil {
			logger.Error("Failed to subscribe to head events")
			time.Sleep(1 * time.Second)
		}

		logger.Warn("Head event SubscribeRaw ended, reconnecting")
	}

}
