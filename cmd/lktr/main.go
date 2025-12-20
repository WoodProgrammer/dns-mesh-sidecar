package main

import (
	"fmt"
	"log"

	"lktr/internal/config"
	"lktr/internal/dns"
	"lktr/internal/server"
	"lktr/pkg/matcher"
)

func main() {
	cfg := config.Load()

	fmt.Printf("DNS Proxy v1.0.0\n")
	fmt.Printf("Listening on: %s\n", cfg.ListenAddr)
	fmt.Printf("Upstream DNS: %s\n", cfg.UpstreamDNS)
	fmt.Println("Starting DNS proxy...")

	blocklist := []string{
		"example.com",
		"*.badsite.com",
		"*.paypal.com",
		"paypal.com",
	}

	m := matcher.BuildMatcher(blocklist)

	dnsHandler := dns.NewHandler(cfg.UpstreamDNS, cfg.Verbose, m)

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
