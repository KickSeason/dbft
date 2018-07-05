package dbft

import (
	"log"
)

type DBFT struct {
	id uint64
}

func NewDBFT(id uint64) *DBFT {
	log.Println("New DBFT.")
	return &DBFT{
		id: id,
	}
}

func HandleMessage() {

}

func SendMessage() {

}

func handlePrepareRequest() {

}

func handlePrepareResponse() {

}

func assemblePrepareRequest() {

}

func assemblePrepareResponse() {

}
