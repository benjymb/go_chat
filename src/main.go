package main

import (
	"log"
	"net/http"
	"views"
)

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", views.HandleConnections)

	// Start listening for incoming models messages
	go views.HandleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
