package server

import (
	"fmt"
	"log"
	"net"

	"lktr/internal/dns"
)

type TCPServer struct {
	ListenAddr string
	Handler    *dns.Handler
	Verbose    bool
}

func NewTCPServer(listenAddr string, handler *dns.Handler, verbose bool) *TCPServer {
	return &TCPServer{
		ListenAddr: listenAddr,
		Handler:    handler,
		Verbose:    verbose,
	}
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on TCP %s: %w", s.ListenAddr, err)
	}
	defer listener.Close()

	log.Printf("DNS proxy listening on TCP %s\n", s.ListenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting TCP connection: %v", err)
			continue
		}

		go s.Handler.HandleTCP(conn)
	}
}
