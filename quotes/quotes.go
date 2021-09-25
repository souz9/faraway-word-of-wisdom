package quotes

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Quotes is a set of quotes.
type Quotes []string

// Load quotes from a file, a quote per line.
func (qs *Quotes) Load(path string) error {
	*qs = Quotes{}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) > 0 {
			*qs = append(*qs, line)
		}
	}
	if err := s.Err(); err != nil {
		return err
	}

	if len(*qs) == 0 {
		return fmt.Errorf("at least one quote is required (input file probably is empty)")
	}
	return nil
}

// Any returns one random quote from the set.
// It panics if the set is empty.
func (qs Quotes) Any() string {
	if len(qs) == 0 {
		panic("at least one quote is required")
	}
	return qs[rand.Intn(len(qs))]
}
