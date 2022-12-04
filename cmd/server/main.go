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
		msg, err := read(conn)
		if err != nil {
			break
		}

		err = write(conn, msg)
		if err != nil {
			break
		}
	}

}

func write(conn net.Conn, msg []byte) error {
	_, err := conn.Write([]byte("From client: " + string(msg)))
	if err != nil {
		log.Println("Error writing client:", err)
		return err
	}

	return nil
}

func read(conn net.Conn) ([]byte, error) {
	msg := make([]byte, 1024)
	_, err := conn.Read(msg)
	if err != nil {
		log.Println("Error reading message:", err)
		return msg, err
	}

	log.Println("Message from client:", string(msg))

	return msg, nil
}
