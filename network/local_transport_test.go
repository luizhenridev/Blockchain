package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	is := assert.New(t)

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	is.Equal(tra.peers[trb.addr], trb)
	is.Equal(trb.peers[tra.addr], tra)

}

func TestSendMessage(t *testing.T) {
	is := assert.New(t)

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello World")

	is.Nil(tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	is.Equal(rpc.Payload, msg)
	is.Equal(rpc.From, tra.addr)

}
