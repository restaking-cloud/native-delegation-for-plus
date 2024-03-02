package k2

import (
	"encoding/json"
	"net/http"
	"strings"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Router paths
	pathRoot                   = "/"
	pathExit                   = "/eth/v1/exit"
	pathClaim                  = "/eth/v1/claim"
	pathRegister               = "/eth/v1/register"
	pathGetDelegatedValidators = "/eth/v1/delegated-validators"
)

func (k2 *K2Service) handleRoot(w http.ResponseWriter, _ *http.Request) {
	k2.respondOK(w, "K2 module is running")
}

func (k2 *K2Service) handleExit(w http.ResponseWriter, r *http.Request) {
	// Post call.
	// Handles the removal of the validators delegated balance from the K2 contract,
	// and therefore exits the K2 protocol.

	payload := phase0.BLSPubKey{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		k2.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := k2.processExit(payload)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	k2.respondOK(w, result)

}

func (k2 *K2Service) handleClaim(w http.ResponseWriter, r *http.Request) {
	// Post call.
	// Handles the claim of rewards for the validators in the K2 contract.

	type claimPayload struct {
		NodeOperators []common.Address `json:"nodeOperators"`
	}

	payload := claimPayload{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		k2.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// If no node operators are provided, it will claim rewards for all the configured node operators.
	if len(payload.NodeOperators) == 0 {
		for _, wallet := range k2.cfg.ValidatorWallets {
			payload.NodeOperators = append(payload.NodeOperators, wallet.Address)
		}
	}

	result, err := k2.batchProcessClaims(payload.NodeOperators)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(result) == 0 {
		// force return an empty array instead of null
		k2.respondOK(w, []string{})
		return
	}

	k2.respondOK(w, result)

}

func (k2 *K2Service) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Post call.
	// Handles the native delegation of validators in the K2 contract if configured to do so.
	// This performs the registration of validators if required, in the Proposer Registry,
	// and then natively delegates them to the K2 contract.
	// If the module is not configured to perform native delegation, it will only register
	// the validators in the Proposer Registry.

	payload := []apiv1.SignedValidatorRegistration{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&payload)
	if err != nil {
		k2.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := k2.batchProcessValidatorRegistrations(payload)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	k2.respondOK(w, result)
}

func (k2 *K2Service) handleGetValidators(w http.ResponseWriter, r *http.Request) {
	// Get call.
	// Handles the retrieval of the delegated validators based on representative addresses.
	// If no representative addresses are provided, it will return the delegated validators
	// for all the configured representative addresses.

	queryParams := r.URL.Query()

	representativeAddresses := []common.Address{}
	if representativeAddressesStr := queryParams.Get("representativeAddresses"); representativeAddressesStr != "" {
		representativeAddressesStr := strings.Split(representativeAddressesStr, ",")
		for _, address := range representativeAddressesStr {
			representativeAddresses = append(representativeAddresses, common.HexToAddress(address))
		}
	} else {
		for _, wallet := range k2.cfg.ValidatorWallets {
			representativeAddresses = append(representativeAddresses, wallet.Address)
		}
	}

	includeBalance := false
	if includeBalanceStr := queryParams.Get("includeBalance"); includeBalanceStr != "" {
		includeBalance = includeBalanceStr == "true"
	}

	result, err := k2.getDelegatedValidators(representativeAddresses, includeBalance)
	if err != nil {
		k2.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(result) == 0 {
		// force return an empty array instead of null
		k2.respondOK(w, []string{})
		return
	}

	k2.respondOK(w, result)
}
