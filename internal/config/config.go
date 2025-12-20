package config

import (
	"flag"
)

type Config struct {
	ListenAddr  string
	UpstreamDNS string
	Verbose     bool
	Blocklist   []string
}

func Load() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.ListenAddr, "listen", ":53", "Address to listen on (default :53)")
	flag.StringVar(&cfg.UpstreamDNS, "upstream", "1.1.1.1:53", "Upstream DNS server (default 1.1.1.1:53)")
	flag.BoolVar(&cfg.Verbose, "verbose", false, "Enable verbose logging")
	flag.Parse()

	return cfg
}
