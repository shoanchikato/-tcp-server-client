package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalln("Error creating listener", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			break
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		msg := make([]byte, 1024)
		_, err := conn.Read(msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Println("Message from client:", string(msg))

		_, err = conn.Write([]byte("From client: " + string(msg)))
		if err != nil {
			log.Println("Error writing client:", err)
			break
		}
	}

}
