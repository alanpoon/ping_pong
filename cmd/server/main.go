package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:5555"
	log.Printf("Ping-pong server listen on %v", addr)
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("%v", err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("%v", err)
		}
		if conn == nil {
			log.Fatal("could not create connection")
		}

		log.Printf("Connection accepted: %v", conn.RemoteAddr())

		go func(conn net.Conn) {
			defer conn.Close()
			rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
			log.Printf("Handling %v", conn.RemoteAddr())
			for {
				req, err := rw.ReadString('\n')
				if err != nil {
					log.Printf("failed to read input: %v", err)
					return
				}
				if !bytes.Equal([]byte(req), []byte{112, 105, 110, 103, 13, 10}) { //ping
					log.Printf("wrong data: %v: [%v]", []byte(req), req)
					return

				}
				log.Printf("Received: %v : %q. OK", []byte(req), req)
				rw.WriteString(string([]byte{112, 111, 110, 103, 13, 10})) //pong
				rw.Flush()

			}
		}(conn)
	}
}
