package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type Message struct {
	Text     string
	Latitude float64
	Long     float64
	Expiry   time.Time
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=goserver dbname=messageinabottle sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
}

// Read text, latitude and long from request
// Set expiry to 24 hours from now
func sendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Send Handler")
}

// Return nearest messages that haven't expired, with all data
func messagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message Handler")

	rows, err := db.Query("SELECT * FROM message")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var text, point string
		var time time.Time
		err := rows.Scan(&text, &point, &time)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(text, point, time)
	}
}

func main() {
	http.HandleFunc("/send/", sendHandler)
	http.HandleFunc("/messages/", messagesHandler)
	http.ListenAndServe(":8080", nil)
}
