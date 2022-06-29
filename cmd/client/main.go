package main

import (
	"bufio"
	"bytes"
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
		rw.WriteString(string([]byte{112, 105, 110, 103, 13, 10})) //ping
		rw.Flush()

		log.Print("Reaging responce")
		resp, err := rw.ReadString('\n')
		if err != nil || !bytes.Equal([]byte(resp), []byte{112, 111, 110, 103, 13, 10}) {
			log.Print("failed to read conn: %v", err)
			return
		}

		log.Printf("Received: %v: %q. OK", []byte(resp), resp)
		time.Sleep(time.Second * 1)
	}
}
