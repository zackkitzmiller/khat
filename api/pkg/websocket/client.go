package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type   int    `json:"type"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type output struct {
	Body   string `json:"message"`
	Author string `json:"username"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf(err.Error())
			return
		}
		incomming := output{}
		json.Unmarshal(p, &incomming)
		message := Message{
			Type:   messageType,
			Body:   incomming.Body,
			Author: incomming.Author,
		}
		c.Pool.Broadcast <- message
	}
}
