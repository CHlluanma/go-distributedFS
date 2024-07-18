package p2p

import "errors"

// ErrInvalidHandshake is returned if the handshake between
// the local and remote node could not be established
var ErrInvalidHandshake = errors.New("InvalidHandshake Error")

// HandshakeFunc ...
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc() HandshakeFunc {
	return func(peer Peer) error {
		return nil
	}
}
