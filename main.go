package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("server running on port: ", *listenAddr)
	log.Fatal(server.Start())
}
