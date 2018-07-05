package main

import (
	"log"
	"dbft"
)

type node struct{
	id uint64
}

func main() {
	log.Println("main run")
	dbft.NewDBFT(0)
}

func run() {
	
}