package dbft

import (
	"log"
)

type Message struct {

}

type MessageQue struct {
	que []Message
}

func NewMessageQue() *MessageQue {
	log.Println("New MessageQue")
	return &MessageQue{
		que: []Message{},
	}
}

func Messages(mq *MessageQue) []Message {
	return []Message{}
}

func AddMessage(mq *MessageQue, megs []Message) {

}