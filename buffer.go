package dbft

import (
	"sync"
)

type Buffer struct {
	lock sync.RWMutex
	data []*Transaction
}

func NewBuffer() *Buffer {
	return &Buffer{
		data: make([]*Transaction, 0, 1024),
	}
}

func (b* Buffer) len() uint64 {
	return uint64(len(b.data))
}

func (b* Buffer) push(tx *Transaction) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.data = append(b.data, tx)
}

func (b* Buffer) get(count uint64) []*Transaction {
	if b.len() < count {
		return b.data[:]
	}
	output := b.data[:count]
	return output
}

func (b* Buffer) delete(txs []*Transaction) {
	b.lock.Lock()
	defer b.lock.Unlock()
	tmp := make(map[string]*Transaction)
	for _, value := range b.data {
		tmp[string(value.Hash())] = value
	}
	for _, value := range txs {
		delete(tmp, string(value.Hash()))
	}
	data := make([]*Transaction, len(tmp))
	for _, value := range tmp {
		data = append(data, value)
	}
	b.data = data
}