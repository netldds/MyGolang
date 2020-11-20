package main

import (
	"crypto/rand"
	"io"
	"log"
	"net"
	"time"
)

func reader(i io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := rand.Read(buf)
		if err != nil {
			return
		}
		log.Println("Client got:", string(buf[0:n]))
	}
}
func main() {
	c, err := net.Dial("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()
	for {
		msg := "hi"
		_, err = c.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		log.Println("Client sent:", msg)
		time.Sleep(1e9)
	}
}
