package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"

	"github.com/bwesterb/go-pow"
	"github.com/souz9/faraway-word-of-wisdom.git/protocol"
	"github.com/souz9/faraway-word-of-wisdom.git/quotes"
)

type Server struct {
	Quotes        quotes.Quotes
	POWDifficulty uint32
}

func (s *Server) Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		// TODO: limit concurrent connections to prevent connection flood.
		// TODO: call SetDeadline to prevent slow read/write attack.

		go func() {
			if err := s.handle(conn); err != nil {
				// Debug print or send the error back to the client
				// if it should need.
				// Here we can ignore any further errors, 'cause most likely the client
				// has already disconnected.
				// conn.Write([]byte(err.Error()))
			}
			conn.Close()
		}()
	}
}

func (s *Server) handle(conn net.Conn) error {
	r, w := bufio.NewReader(conn), conn

	// Start of POW-procedure.
	// Generate and send "challenge" to the client.
	challenge := pow.NewRequest(s.POWDifficulty, nonce())
	if err := protocol.Write(w, challenge); err != nil {
		return err
	}

	// Wait for response (POW solution) and verify it.
	proof, err := protocol.Read(r)
	if err != nil {
		return err
	}
	verified, err := pow.Check(challenge, proof, nil)
	if err != nil || !verified {
		return fmt.Errorf("access denied: %v", err)
	}
	// End of POW-procedure.

	return protocol.Write(w, s.Quotes.Any())
}

// Generates random sequence that's used as a nonce.
func nonce() []byte {
	nonce := make([]byte, 8)
	rand.Read(nonce)
	return nonce
}
