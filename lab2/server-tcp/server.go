package main

import (
	"log"
	"net"

	"handouts/lab2/kvs"
)

var keyVal = kvs.NewKvs()

func main() {
	if err := run(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run() error {
	return nil
}

func handleConn(conn net.Conn) {
}
