package main

import (
	"dbft"
	"dbft/network"
)

const (
	nodecnt = 7
	faultcnt = 2
	batchSize = 500
)

func main() {
	nodes := makeNetwork(nodecnt)
	msg := dbft.Message{
		From: 1,
		To: 0,
		Payload: nil,
	}
	nodes[0].Tr.Receiver() <- msg
	for {
		
	}
}

func makeNetwork(n int) [] *network.Node {
	nodes := make([] *network.Node, nodecnt)
	for i := 0; i < n; i++ {
		nodes[i] = network.NewNode(uint64(i))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			nodes[j].AddEndpoit(uint64(j), nodes[j].Receiver())
		}
	}

	for i := 0; i < n; i++ {
		nodes[i].Start()
	}
	return nodes
}