package main

import (
	"bufio"
	"net"

	"github.com/bwesterb/go-pow"
	"github.com/souz9/faraway-word-of-wisdom.git/protocol"
)

type Client struct {
	Hostport string // host:port to a words-of-wisdom server.
}

// Makes a request to the server.
// Returns a quote.
func (c Client) GetQuote() (string, error) {
	conn, err := net.Dial("tcp", c.Hostport)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	r, w := bufio.NewReader(conn), conn

	// Start of POW-procedure.
	// Receive a challenge and fulfill the proof.
	challenge, err := protocol.Read(r)
	if err != nil {
		return "", err
	}
	proof, err := pow.Fulfil(challenge, nil)
	if err != nil {
		return "", err
	}
	if err := protocol.Write(w, proof); err != nil {
		return "", err
	}
	// End of POW-procedure.

	return protocol.Read(r)
}
