package protocol

import (
	"bufio"
	"io"
	"strings"
)

// Reads a message from connection.
func Read(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	return strings.TrimSuffix(line, "\n"), err
}

// Writes a message to connection.
func Write(w io.Writer, message string) error {
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
