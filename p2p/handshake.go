package p2p

// HandshakeFunc ...
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc() HandshakeFunc {
	return func(peer Peer) error {
		return nil
	}
}
