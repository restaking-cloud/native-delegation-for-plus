package beacon

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/restaking-cloud/native-delegation-for-plus/beacon/config"
)

type BeaconService struct {
	cfg    config.BeaconConfig
	client *http.Client
}

func NewBeaconService() *BeaconService {
	return &BeaconService{
		client: &http.Client{Timeout: 6 * time.Second}, // get a response in half a slot
	}
}

func (b *BeaconService) Configure(cfg config.BeaconConfig) error {
	b.cfg = cfg

	err := b.connect(cfg.BeaconNodeUrl)
	if err != nil {
		return err
	}

	return nil
}

func (b *BeaconService) ConnectedChainId() *big.Int {
	return b.cfg.ChainID
}

func (b *BeaconService) Status() (res *SyncStatusData, err error) {
	return b.syncProgress(context.Background())
}
