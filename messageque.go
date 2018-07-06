package dbft

import (
	"sync"
)

type Message struct {
	From	uint64
	To		uint64
	Payload interface{} 	
}

type MessageQue struct {
	lock sync.RWMutex
	que []*Message
}

func NewMessageQue() *MessageQue {
	return &MessageQue{
		que: []*Message{},
	}
}

func (mq *MessageQue)len() uint64 {
	return uint64(len(mq.que))
}

func (mq* MessageQue)Push(m *Message) {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	mq.que = append(mq.que, m)
}

func (mq* MessageQue)PopN(count uint64) []*Message {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if count < mq.len() {
		output := mq.que
		mq.que = make([]*Message, 0, 1024)
		return output
	}
	output := mq.que[:count]
	mq.que = mq.que[count:]
	return output	
}

func (mq* MessageQue)PopOne() *Message {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if mq.len() < 1 {
		return nil
	}
	output := mq.que[0]
	if mq.len() == 1 {
		mq.que = make([]*Message, 0, 1024)
	} else {
		mq.que = mq.que[1:]
	}
	return output
}