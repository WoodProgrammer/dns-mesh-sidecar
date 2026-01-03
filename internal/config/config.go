package config

import (
	"flag"
	"time"
)

type Config struct {
	ListenAddr            string
	UpstreamDNS           string
	Verbose               bool
	Blocklist             []string
	DryRun                bool
	ControllerURL         string
	FetchInterval         time.Duration
	MetricsAddr           string
	HTTPSModeEnabled      bool
	HTTPSUpstream         string
	TLSCACert             string
	TLSClientCert         string
	TLSClientKey          string
	TLSInsecureSkipVerify bool
}

func Load() *Config {
	cfg := &Config{}
	fetchIntervalSec := 0

	flag.StringVar(&cfg.ListenAddr, "listen", ":53", "Address to listen on (default :53)")
	flag.StringVar(&cfg.UpstreamDNS, "upstream", "1.1.1.1:53", "Upstream DNS server (default 1.1.1.1:53)")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Enable verbose logging")
	flag.StringVar(&cfg.ControllerURL, "controller", "", "Controller URL to fetch policies from")
	flag.IntVar(&fetchIntervalSec, "fetch-interval", 30, "Policy fetch interval in seconds (default 30)")
	flag.StringVar(&cfg.MetricsAddr, "metrics", ":9090", "Metrics HTTP server address (default :9090)")
	flag.BoolVar(&cfg.HTTPSModeEnabled, "https-mode", false, "Enable DNS-over-HTTPS mode")
	flag.StringVar(&cfg.HTTPSUpstream, "https-upstream", "https://1.1.1.1/dns-query", "DNS-over-HTTPS upstream server (default Cloudflare)")
	flag.StringVar(&cfg.TLSCACert, "tls-ca-cert", "", "Path to CA certificate for verifying DoH server")
	flag.StringVar(&cfg.TLSClientCert, "tls-client-cert", "", "Path to client certificate for mTLS")
	flag.StringVar(&cfg.TLSClientKey, "tls-client-key", "", "Path to client private key for mTLS")
	flag.BoolVar(&cfg.TLSInsecureSkipVerify, "tls-insecure-skip-verify", false, "Skip TLS certificate verification (insecure, for testing only)")
	flag.Parse()

	cfg.FetchInterval = time.Duration(fetchIntervalSec) * time.Second

	return cfg
}
