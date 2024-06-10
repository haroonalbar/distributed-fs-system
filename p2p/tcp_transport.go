package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport (listenAddr string) *TCPTransport {
  return &TCPTransport{
    listenAddress: listenAddr,
  }
}
