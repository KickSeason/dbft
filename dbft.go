package dbft

import (
	"log"
	"errors"
)

type DBFTMessage struct {

	payload interface{}
}

type prepareRequest struct {
	pid uint64
	vid uint64
	txs []Transaction
}

type prepareResponse struct {
	pid uint64
	vid	uint64
	txs []Transaction
}

type DBFT struct {
	id uint64
	buffer []*Transaction
	mq *MessageQue
}

func NewDBFT(id uint64) *DBFT {
	return &DBFT{
		id: id,
		mq: NewMessageQue(),
	}
}

func (d *DBFT) HandleMessage(sendID uint64, msg *DBFTMessage) error {
	switch t := msg.payload.(type) {
	case *prepareRequest:
		return d.handlePrepareRequest(sendID, t)
	case *prepareResponse:
		return d.handlePrepareResponse(sendID, t)
	default:
		log.Println("unknow dbft message type.")
		return errors.New("unknow dbft message type.")
	}
}

func (d *DBFT) SendMessage() {

}

func (d *DBFT) handlePrepareRequest(sendID uint64, msg *prepareRequest) error {
	return nil
}

func (d *DBFT) handlePrepareResponse(sendID uint64, msg *prepareResponse) error {
	return nil
}

func (d *DBFT) assemblePrepareRequest() {

}

func (d *DBFT) assemblePrepareResponse() {

}
