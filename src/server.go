package main

import (
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Text     string
	Latitude float64
	Long     float64
	Expiry   time.Time
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Send Handler")
	// Read text, latitude and long from request
	// Set expiry to 24 hours from now
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message Handler")
	// Return nearest messages that haven't expired, with all data
}

func main() {
	http.HandleFunc("/send/", sendHandler)
	http.HandleFunc("/messages/", messagesHandler)
	http.ListenAndServe(":8080", nil)
}
