package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strings"

	"github.com/bwesterb/go-pow"
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
				// The client probably will be interested in the error if occurred,
				// trying to send it as a text.
				// Here we can ignore any further errors, 'cause most likely the client
				// has already disconnected.
				conn.Write([]byte(err.Error()))
			}
			conn.Close()
		}()
	}
}

func (s *Server) handle(conn net.Conn) error {
	r := bufio.NewReader(conn)
	w := conn

	// Start of POW-procedure.
	// Generate and send "challenge" to the client.
	challenge := pow.NewRequest(s.POWDifficulty, nonce())
	if err := s.write(w, challenge); err != nil {
		return err
	}

	// Wait for response (POW solution) and verify it.
	proof, err := s.read(r)
	if err != nil {
		return err
	}
	verified, err := pow.Check(challenge, proof, nil)
	if err != nil || !verified {
		return fmt.Errorf("access denied (%v)", err)
	}
	// End of POW-procedure.

	return s.write(w, s.Quotes.Any())
}

// Reads a message from connection.
func (_ *Server) read(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	return strings.TrimSuffix(line, "\n"), err
}

// Writes a message to connection.
func (_ *Server) write(w io.Writer, message string) error {
	line := append([]byte(message), '\n')
	for len(line) > 0 {
		n, err := w.Write(line)
		if err != nil {
			return err
		}
		line = line[n:]
	}
	return nil
}

// Generates random sequence that's used as a nonce.
func nonce() []byte {
	nonce := make([]byte, 8)
	rand.Read(nonce)
	return nonce
}
