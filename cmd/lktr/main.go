package main

import (
	"fmt"
	"log"

	"lktr/internal/client"
	"lktr/internal/config"
	"lktr/internal/dns"
	"lktr/internal/server"
	"lktr/pkg/matcher"
)

func main() {
	cfg := config.Load()

	fmt.Printf("DNS Proxy v1.0.0 (Sidecar Mode)\n")
	fmt.Printf("Listening on: %s\n", cfg.ListenAddr)
	fmt.Printf("Upstream DNS: %s\n", cfg.UpstreamDNS)
	if cfg.ControllerURL != "" {
		fmt.Printf("Controller URL: %s\n", cfg.ControllerURL)
		fmt.Printf("Fetch Interval: %v\n", cfg.FetchInterval)
	}
	fmt.Println("Starting DNS proxy...")

	blocklist := []string{}

	m := matcher.BuildMatcher(blocklist)

	dnsHandler := dns.NewHandler(cfg.UpstreamDNS, cfg.Verbose, m)

	updateChannel := make(chan []string, 10)

	go func() {
		for newBlocklist := range updateChannel {
			if cfg.Verbose {
				log.Printf("Received blocklist update with %d entries", len(newBlocklist))
			}

			newMatcher := matcher.BuildMatcher(newBlocklist)
			dnsHandler.UpdateMatcher(newMatcher)

			fmt.Printf("Blocklist updated successfully with %d entries\n", len(newBlocklist))
		}
	}()

	if cfg.ControllerURL != "" {
		fetcher := client.NewFetcher(cfg.ControllerURL, cfg.FetchInterval, cfg.Verbose, updateChannel)
		go fetcher.Start()
	} else {
		log.Println("Warning: No controller URL specified, running without policy updates")
	}

	udpServer := server.NewUDPServer(cfg.ListenAddr, dnsHandler, cfg.Verbose)
	tcpServer := server.NewTCPServer(cfg.ListenAddr, dnsHandler, cfg.Verbose)

	go func() {
		if err := udpServer.Start(); err != nil {
			log.Fatalf("UDP server error: %v", err)
		}
	}()

	if err := tcpServer.Start(); err != nil {
		log.Fatalf("TCP server error: %v", err)
	}
}
