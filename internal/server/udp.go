package server

import (
	"fmt"
	"log"
	"net"

	"lktr/internal/dns"
)

type UDPServer struct {
	ListenAddr string
	Handler    *dns.Handler
	Verbose    bool
}

func NewUDPServer(listenAddr string, handler *dns.Handler, verbose bool) *UDPServer {
	return &UDPServer{
		ListenAddr: listenAddr,
		Handler:    handler,
		Verbose:    verbose,
	}
}

func (s *UDPServer) Start() error {
	addr, err := net.ResolveUDPAddr("udp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP %s: %w", s.ListenAddr, err)
	}
	defer conn.Close()

	fmt.Printf("DNS proxy listening on UDP %s\n", s.ListenAddr)

	buffer := make([]byte, 512)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		if s.Verbose {
			log.Printf("Received %d bytes from %s", n, clientAddr)
		}

		queryCopy := make([]byte, n)
		copy(queryCopy, buffer[:n])

		go s.Handler.HandleUDP(conn, clientAddr, queryCopy)
	}
}
