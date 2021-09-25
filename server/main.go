package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/souz9/faraway-word-of-wisdom.git/quotes"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	listen := flag.String("listen", ":9000", "Listen incoming requests on host:port")
	quotesPath := flag.String("quotes", "quotes.txt", "Path to a file with quotes")
	flag.Parse()

	var qs quotes.Quotes
	if err := qs.Load(*quotesPath); err != nil {
		log.Printf("load quotes: %v", err)
		os.Exit(1)
	}

	s := Server{
		Quotes:        qs,
		POWDifficulty: 10,
	}
	log.Printf("start listening incoming requests on %q", *listen)
	if err := s.Listen(*listen); err != nil {
		log.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}
