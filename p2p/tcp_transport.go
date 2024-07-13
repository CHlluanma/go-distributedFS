package p2p

import (
	"bytes"
	"fmt"

	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established conection
type TCPPeer struct {
	// conn is the underlying conncetion of the peer
	conn net.Conn

	// if we dial and retrieve a conn -> oubound == true
	// if we accept and retrieve a conn -> oubound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddr    string
	listener      net.Listener
	handshakeFunc HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr:    listenAddr,
		handshakeFunc: NOPHandshakeFunc(),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.handshakeFunc(peer); err != nil {

	}

	// Read Loop
	buf := new(bytes.Buffer)
	for {
		if err := t.decoder.Decode(conn, buf); err != nil {
			fmt.Printf("TCP decode error: %s\n", err)
			continue
		}
	}
}
