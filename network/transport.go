package network

import (
	"dbft"
)

type Transport struct {
	receiver 	chan dbft.Message
	sender		chan dbft.Message
	addr		uint64
	endpoints 	map[uint64] chan<- dbft.Message
}

func NewTransport(id uint64) *Transport {
	return &Transport{
		addr: id,
		receiver: 	make(chan dbft.Message),
		sender: 	make(chan dbft.Message),
		endpoints: 	map[uint64] chan<- dbft.Message {},
	}
}

func (t *Transport)SendMessage( msg dbft.Message) {
	t.endpoints[msg.To]<- msg
}

func (t *Transport)AddEndpoit(addr uint64, out chan<- dbft.Message) {
	t.endpoints[addr] = out
}

func (t *Transport)Receiver() chan dbft.Message {
	return t.receiver
}

func (t *Transport)Sender() chan dbft.Message {
	return t.sender
}
