package main

import (
	"h2/h2"
	"log"
)

func main() {
	log.Println("server has started")
	h2.StartServer()
}
