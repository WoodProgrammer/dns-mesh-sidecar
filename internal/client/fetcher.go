package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type PolicyResponse struct {
	Blocklist []string `json:"blocklist"`
}

type Fetcher struct {
	controllerURL string
	fetchInterval time.Duration
	verbose       bool
	updateChannel chan []string
	httpClient    *http.Client
}

func NewFetcher(controllerURL string, fetchInterval time.Duration, verbose bool, updateChannel chan []string) *Fetcher {
	return &Fetcher{
		controllerURL: controllerURL,
		fetchInterval: fetchInterval,
		verbose:       verbose,
		updateChannel: updateChannel,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (f *Fetcher) Start() {
	if f.verbose {
		log.Printf("Starting policy fetcher, controller: %s, interval: %v", f.controllerURL, f.fetchInterval)
	}

	ticker := time.NewTicker(f.fetchInterval)
	defer ticker.Stop()

	// Fetch immediately on start
	f.fetchPolicies()

	for range ticker.C {
		f.fetchPolicies()
	}
}

func (f *Fetcher) fetchPolicies() {
	if f.verbose {
		log.Printf("Fetching policies from controller: %s", f.controllerURL)
	}

	resp, err := f.httpClient.Get(f.controllerURL)
	if err != nil {
		log.Printf("Error fetching policies: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from controller: %d", resp.StatusCode)
		return
	}

	var policyResp PolicyResponse
	if err := json.NewDecoder(resp.Body).Decode(&policyResp); err != nil {
		log.Printf("Error decoding policy response: %v", err)
		return
	}

	if f.verbose {
		log.Printf("Fetched %d policy entries from controller", len(policyResp.Blocklist))
	}

	f.updateChannel <- policyResp.Blocklist

	fmt.Printf("Policies fetched successfully: %d entries\n", len(policyResp.Blocklist))
}
