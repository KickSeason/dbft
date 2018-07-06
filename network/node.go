package network

import (
	"log"
	"dbft"
	"errors"
)

type Node struct{
	id 		uint64
	runing 	bool
	dbft 	*dbft.DBFT
	Tr 		*Transport
}

func NewNode(id uint64) *Node {
	return &Node{
		id: id,
		runing: false,
		dbft: dbft.NewDBFT(id),
		Tr: NewTransport(id),
	}
}

func (n* Node)Start() {
	n.runing = true
	go n.run()
}

func (n* Node)run() {
	for n.runing {
		select {
		case msg := <- n.Tr.Receiver():
			n.handleMessage(&msg)
		case msg := <- n.Tr.Sender():
			n.Tr.SendMessage(msg)
		}
	}
}

func (n* Node)handleMessage(msg *dbft.Message) error {
	switch t := msg.Payload.(type) {
		case *dbft.DBFTMessage:
			return n.dbft.HandleMessage(msg.From, t)
		default:
			log.Println("[node](HandleMessage) unknow message type.")
			return errors.New("unknow message type.")
	}
}
func (n* Node)Stop() {
	n.runing = false
}

func (n *Node)AddEndpoit(addr uint64, out chan<- dbft.Message) {
	n.Tr.AddEndpoit(addr, out)
}

func (n *Node)Receiver() chan<- dbft.Message {
	return n.Tr.Receiver()
}
