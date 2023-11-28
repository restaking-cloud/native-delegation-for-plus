package web3signer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"io"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
)

type Web3SignerService struct {
	url    *url.URL
	client *http.Client
}

func NewWeb3SignerService() *Web3SignerService {
	return &Web3SignerService{
		client: &http.Client{Timeout: 6 * time.Second}, // get a response in half a slot
	}
}

func (s *Web3SignerService) Configure(url *url.URL) error {
	s.url = url

	health, err := s.GetHealth()
	if err != nil {
		return err
	}

	if health != "UP" {
		return fmt.Errorf("web3signer healthcheck failed: %s", health)
	}

	return nil
}

func (s *Web3SignerService) Status() (string, error) {
	health, err := s.GetHealth()
	if err != nil {
		return "", err
	}

	if health != "UP" {
		return "", fmt.Errorf("web3signer healthcheck failed: %s", health)
	}

	return health, nil
}

func (s *Web3SignerService) GetUrl() *url.URL {
	return s.url
}

func (s *Web3SignerService) SignRegistration(
	FeeRecipient bellatrix.ExecutionAddress,
	Domain uint64,
	Pubkey phase0.BLSPubKey,
	Timestamp time.Time,
) (apiv1.SignedValidatorRegistration, error) {

	payload := ValidatorRegistrationPayload{
		Type: VALIDATOR_REGISTRATION_ACTION,
		ValidatorRegistration: &apiv1.ValidatorRegistration{
			FeeRecipient: FeeRecipient,
			GasLimit:     Domain,
			Pubkey:       Pubkey,
			Timestamp:    Timestamp,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return apiv1.SignedValidatorRegistration{}, err
	}
	signingPath := fmt.Sprintf("%s%s", SignPath, Pubkey.String())

	req, err := http.NewRequestWithContext(context.Background(), "POST", s.url.String()+signingPath, bytes.NewReader(payloadBytes))
	if err != nil {
		return apiv1.SignedValidatorRegistration{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return apiv1.SignedValidatorRegistration{}, err
	}

	if resp.StatusCode != 200 {
		return apiv1.SignedValidatorRegistration{}, fmt.Errorf("invalid response (%d): %v", resp.StatusCode, resp.Body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return apiv1.SignedValidatorRegistration{}, fmt.Errorf("error reading response body: %v", err)
	}
	var phase0Signature phase0.BLSSignature
	copy(phase0Signature[:], body[:])

	res := apiv1.SignedValidatorRegistration{
		Message: payload.ValidatorRegistration,
		Signature: phase0Signature,
	}

	return res, nil
}

func (s *Web3SignerService) GetPubkeyList() (map[string]bool, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", s.url.String()+ListBLSPubKeysPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid response (%d): %v", resp.StatusCode, resp)
	}

	var response []string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	pubkeys := make(map[string]bool)
	for _, pubkey := range response {
		pubkeys[pubkey] = true
	}

	return pubkeys, nil
}

func (s *Web3SignerService) GetHealth() (string, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", s.url.String()+HealthCheckPath, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("invalid response (%d): %v", resp.StatusCode, resp)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response["outcome"].(string), nil
}
