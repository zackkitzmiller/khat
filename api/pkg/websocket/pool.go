package websocket

import "log"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("Size of Connection Pool: ", len(p.Clients))
			for client := range p.Clients {
				log.Printf("%+v", client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user joined..."})
			}
			break
		case client := <-p.Unregister:
			delete(p.Clients, client)
			log.Println("Size of Connection Pool: ", len(p.Clients))
			for client := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-p.Broadcast:
			log.Println("Sending message to all clients")
			for client := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
}
