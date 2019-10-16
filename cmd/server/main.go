package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8080"
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
				if req != "ping\n" {
					log.Printf("wrong data: %v: [%v]", []byte(req), req)
					return

				}
				log.Printf("Received: %v : %q. OK", []byte(req), req)
				rw.WriteString("pong\n")
				rw.Flush()

			}
		}(conn)
	}
}
