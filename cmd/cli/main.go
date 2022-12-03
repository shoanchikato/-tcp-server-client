package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalln("Error dialing:", err)
	}

	for {
		fmt.Printf("Enter data: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		data := scanner.Text()

		if len(data) == 0 {
			continue
		}

		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Fatalln("Error writing to connection:", err)
			break
		}

		msg := make([]byte, 1024)
		_, err = conn.Read(msg)
		if err != nil {
			log.Fatalln("Error reading from connection:", err)
			break
		}

		log.Println("Message from server:", string(msg))
	}
}
