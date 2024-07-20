package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ahang7/go-distributedFS/p2p"
)

func makeServer(listenAddr, root string, nodes ...string) *FileServer {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc(),
		Decoder:       p2p.DefaultDecoder{},
		OnPeer: func(p p2p.Peer) error {
			fmt.Printf("OnPeer: %+v", p)
			return nil
		},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       root,
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tr,
		BootstrapNetWork:  nodes,
	}

	s := NewFileServer(fileServerOpts)
	tr.OnPeer = s.OnPeer

	return s
}

func main() {
	s := makeServer(":3000", "3000_network", ":4000")

	go func() {
		time.Sleep(time.Second * 5)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
