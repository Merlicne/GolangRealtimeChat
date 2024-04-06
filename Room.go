package main

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type ClientList map[*Client]bool
type ChatHistory []([]byte)

type Room struct {
	room_name   string
	Connections ClientList
	ChatHist    ChatHistory
}

func newRoom(name string) *Room {
	return &Room{
		room_name:   name,
		Connections: make(ClientList),
	}
}

func (s *Room) handleWS(ws *websocket.Conn) {

	fmt.Println("New connection")
	c := newClient(ws)
	s.Connections[c] = true

	s.readLoop(c)
}

// func (s *Room) JoinRoom(c *Client) {
// 	s.Connections[c] = true
// 	s.readLoop(c)
// }

// waiting for message broadcast
func (s *Room) readLoop(c *Client) {
	buf := make([]byte, 1024)
	// read old messages
	s.reChat(c)
	for {
		// read if any message
		n, err := c.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("a user had disconnected ：", err)
				delete(s.Connections, c)
				break
			}
			fmt.Println("read error : ", err)
			continue
		}
		msg := make([]byte, n)
		copy(msg, buf[:n])
		s.ChatHist = append(s.ChatHist, msg)
		s.broadcast(msg)
	}
}

// send message to everyone in the room
func (s *Room) broadcast(b []byte) {
	for c := range s.Connections {
		go func(c *Client) {
			if _, err := c.conn.Write(b); err != nil {
				fmt.Println("Write error : ", err)
			}
		}(c)
	}
}

// get all message in the room
func (s *Room) reChat(c *Client) {
	go func(c *Client) {
		for _, text := range s.ChatHist {
			if _, err := c.conn.Write(text); err != nil {
				fmt.Println("Read old message error : ", err)
			}
		}
	}(c)
}
