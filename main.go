package main

import (
	"fmt"
	"log"

	"github.com/ahang7/go-distributedFS/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("OnPeer", peer)
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc(),
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		msg := <-tr.Consume()
		fmt.Printf("message: %+v", msg)
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
