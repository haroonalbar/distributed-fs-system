package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listnerAddr := ":4000"
	tr := NewTCPTransport(listnerAddr)
	assert.Equal(t, tr.listenAddress, listnerAddr)
}
