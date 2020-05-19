package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zackkitzmiller/khat/pkg/websocket"
)

func main() {
	log.Printf("Khat v0.1")
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/", statusHandler())
	http.HandleFunc("/ws", websocketHandler(pool))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("RUnning on port %s", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func websocketHandler(pool *websocket.Pool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("WebSocket Endpoint Hit")

		conn, err := websocket.Upgrade(w, r)
		if err != nil {
			log.Printf(err.Error())
		}

		client := &websocket.Client{
			Conn: conn,
			Pool: pool,
		}

		pool.Register <- client
		client.Read()
	})
}

func statusHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
		return
	})
}
