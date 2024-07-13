package main

import (
	"fmt"
	"github.com/ahang7/go-distributedFS/p2p"
	"log"
)

func main() {

	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello world")
}
