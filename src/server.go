package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Text      string
	Latitude  float64
	Longitude float64
	Expiry    time.Time
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=goserver dbname=messageinabottle sslmode=disable")
	if err != nil {
		log.Println(err.Error())
	}
}

// Read text, latitude and long from request
// Set expiry to 24 hours from now
func sendHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered Send Handler")

	decoder := json.NewDecoder(r.Body)
	var message Message
	err := decoder.Decode(&message)
	if err != nil {
		log.Println(err.Error())
	}

	t := time.Now()
	t = t.AddDate(0, 0, 1)
	_, err = db.Exec(
		"INSERT INTO message (text, location, expiry) VALUES ($1, point($2, $3), $4)",
		message.Text,
		message.Latitude,
		message.Longitude,
		t,
	)

	if err != nil {
		log.Println(err.Error())
	}

	w.WriteHeader(200)
}

// Return nearest messages that haven't expired, with all data
func messagesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered Messages Handler")

	rows, err := db.Query("SELECT * FROM message")
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		var text, point string
		var time time.Time
		err := rows.Scan(&text, &point, &time)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(text, point, time)
	}
}

func main() {
	http.HandleFunc("/send/", sendHandler)
	http.HandleFunc("/messages/", messagesHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
