package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"testing"

	"github.com/bwesterb/go-pow"
	"github.com/souz9/faraway-word-of-wisdom.git/quotes"
	"github.com/stretchr/testify/require"
)

func TestServerHandler(t *testing.T) {
	srv := &Server{
		Quotes:        quotes.Quotes{"QUOTE"},
		POWDifficulty: 1,
	}

	r, w := fakeConnect(t, srv)
	challenge := readline(t, r)
	proof, err := pow.Fulfil(challenge, nil)
	require.NoError(t, err)
	writeline(t, w, proof)

	q := readline(t, r)
	require.Equal(t, "QUOTE", q)

	t.Run("when a client provides invalid proof", func(t *testing.T) {
		r, w := fakeConnect(t, srv)

		// Ignore challenge, respond with rubbish to the server.
		readline(t, r)
		writeline(t, w, "RUBBISH")

		// Assert that the server responds with nothing.
		res, err := ioutil.ReadAll(r)
		require.NoError(t, err)
		require.Empty(t, res)
	})
}

// Establishes a fake connection to the server.
// Returns client's reader and writer.
func fakeConnect(t *testing.T, s *Server) (*bufio.Reader, io.Writer) {
	a, b := net.Pipe()
	go func() {
		s.handle(b)
		b.Close()
	}()
	return bufio.NewReader(a), a
}

func readline(t *testing.T, r *bufio.Reader) string {
	t.Helper()
	line, err := r.ReadString('\n')
	require.NoError(t, err)
	return strings.TrimSuffix(line, "\n")
}

func writeline(t *testing.T, w io.Writer, s string) {
	t.Helper()
	line := append([]byte(s), '\n')
	n, err := w.Write(line)
	require.NoError(t, err)
	require.Equal(t, len(line), n)
}
