package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	hostport := flag.String("connect", "localhost:9000", "Connection string to a words-of-wisdom server, host:port")
	flag.Parse()

	quote, err := Client{Hostport: *hostport}.GetQuote()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Println(quote)
}
