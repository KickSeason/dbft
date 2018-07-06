package dbft

import (
	"math/rand"
	"encoding/binary"
)

type Transaction struct {
	value uint64
}

func NewTransaction() *Transaction {
	return &Transaction{rand.Uint64()}
}

func (t *Transaction) Hash() []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, t.value)
	return buf
}