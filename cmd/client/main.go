package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:8080"
	log.Printf("Trying to connect to: %v", addr)
	client, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("%v", err)
	}
	log.Printf("Connection established %v", client.RemoteAddr())
	rw := bufio.NewReadWriter(bufio.NewReader(client), bufio.NewWriter(client))

	for {
		log.Print("Sending: \"ping\\n\"")
		rw.WriteString("ping\n")
		rw.Flush()

		log.Print("Reaging responce")
		resp, err := rw.ReadString('\n')
		if err != nil || resp != "pong\n" {
			log.Print("failed to read conn: %v", err)
			return
		}

		log.Printf("Received: %v: %q. OK", []byte(resp), resp)
		time.Sleep(time.Second * 1)
	}
}
