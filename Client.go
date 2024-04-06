package main

import "golang.org/x/net/websocket"

type Client struct {
	conn *websocket.Conn
}

func newClient(c *websocket.Conn) *Client {
	return &Client{
		conn: c,
	}
}
