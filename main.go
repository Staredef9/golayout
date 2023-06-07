package main

import (
	"flag"
	"log"
	"prova"
)

func main() {

	listenAddr := flag.String("listenaddr", ":3000", "the server address")

	server := api.NewServer(*listenAddr)
	log.Fatal(server.Start())

	flag.Parse()
}
