package signatureswapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"

	"github.com/ethereum/go-ethereum/common"

	"github.com/restaking-cloud/native-delegation-for-plus/signatureswapper/config"
)

type SignatureSwapperService struct {
	client *http.Client
	cfg    config.SignatureSwapperConfig
}

func NewSignatureSwapperService() *SignatureSwapperService {
	return &SignatureSwapperService{
		client: &http.Client{Timeout: 6 * time.Second}, // get a response in half a slot
	}
}

func (s *SignatureSwapperService) Configure(url *url.URL) error {
	s.cfg.Url = url

	info, err := s.GetInfo()
	if err != nil {
		return err
	}

	s.cfg.Domain = info.GasLimitProposerRegistryDomain
	s.cfg.ChainID = big.NewInt(int64(info.ChainID))

	return nil
}

func (s *SignatureSwapperService) Domain() uint64 {
	return s.cfg.Domain
}

func (s *SignatureSwapperService) ConnectedChainId() *big.Int {
	return s.cfg.ChainID
}

func (s *SignatureSwapperService) GetInfo() (Info, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", s.cfg.Url.String()+InfoPath, nil)
	if err != nil {
		return Info{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return Info{}, err
	}

	if resp.StatusCode != 200 {
		return Info{}, fmt.Errorf("invalid response (%d): %v", resp.StatusCode, resp)
	}

	var response Info
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Info{}, err
	}

	return response, nil
}

func (s *SignatureSwapperService) GenerateSignature(
	registration apiv1.SignedValidatorRegistration,
	representativeAddress common.Address,
) (EcdsaSignature, error) {

	payload := SignatureSwapPayload{
		Signature:             registration.Signature,
		Message:               registration.Message,
		RepresentativeAddress: representativeAddress,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return EcdsaSignature{}, err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", s.cfg.Url.String()+GenerateSignaturePath, bytes.NewReader(payloadBytes))
	if err != nil {
		return EcdsaSignature{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return EcdsaSignature{}, err
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return EcdsaSignature{}, fmt.Errorf("error reading invalid response (%d) body: %v", resp.StatusCode, err)
		}

		return EcdsaSignature{}, fmt.Errorf("invalid response (%d): %v", resp.StatusCode, string(body))
	}

	var response SignatureSwapResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return EcdsaSignature{}, err
	}

	return response.EcdsaSignature, nil
}
