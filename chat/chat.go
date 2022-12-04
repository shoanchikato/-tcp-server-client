package chat

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Chat interface {
	Command(conn net.Conn, message string)
	// Join(conn net.Conn, username string)
	Leave(conn net.Conn)
	// Broadcast(message string)
	// Control()
}
type userConn struct {
	username string
	conn     net.Conn
}

type chat struct {
	join      chan userConn
	leave     chan userConn
	broadcast chan string
	group     map[net.Conn]string
	write     func(conn net.Conn, message []byte) error
}

func NewChat(write func(conn net.Conn, msg []byte) error) Chat {
	return &chat{
		make(chan userConn),
		make(chan userConn),
		make(chan string),
		make(map[net.Conn]string),
		write,
	}
}

func (c *chat) Control() {
	select {
	case user := <-c.join:
		c.group[user.conn] = user.username
		fmt.Println(len(c.group))
		c.Broadcast(user.username + "joined")
	case user := <-c.leave:
		delete(c.group, user.conn)
		c.Broadcast(user.username + "left")
	case message := <-c.broadcast:
		for conn, _ := range c.group {
			err := c.write(conn, []byte(message))
			log.Fatalln("Error writing in Chat", err)
		}
	default:
	}
}

func (c *chat) Command(conn net.Conn, message string) {
	split := strings.Split(message, " ")
	command := split[0]
	msg := strings.Join(split[1:], " ")

	fmt.Println("command:", command, "message:", strings.Join(split[1:], " "), "msg", msg)

	switch command {
	case "join":
		c.Join(conn, msg)
	case "leave":
		c.Leave(conn)
	case "broadcast":
		c.Broadcast(msg)
	default:
	}

	go c.Control()
}

func (c *chat) Join(conn net.Conn, username string) {
	c.join <- userConn{
		username,
		conn,
	}
}

func (c *chat) Leave(conn net.Conn) {
	username, ok := c.group[conn]
	if ok {
		username := username
		c.leave <- userConn{username, conn}
	}
}

func (c *chat) Broadcast(message string) {
	fmt.Println("message:", message)
	c.broadcast <- message
}
