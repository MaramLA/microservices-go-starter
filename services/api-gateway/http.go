package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 9)
	var reqBody previewTripRequest
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	if reqBody.UserId == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		http.Error(w, "failed to marshal request body", http.StatusInternalServerError)
		return
	}
	reader := bytes.NewReader(jsonBody)

	// TODO: call trip service
	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Print(err)
		http.Error(w, "failed to call trip service", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		log.Println("err: ", err)
		http.Error(w, "failed to parse response from trip service", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: respBody}

	writeJSON(w, http.StatusOK, response)
}
