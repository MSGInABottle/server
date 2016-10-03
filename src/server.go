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

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(200)
}

// Return nearest messages that haven't expired, with all data
func messagesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered Messages Handler")

	var message Message
	messages := []Message{}

	rows, err := db.Query(
		"SELECT text, location[0], location[1], expiry FROM message WHERE expiry >= NOW() AND location <@ circle(point($1, $2), 0.001)",
		r.URL.Query().Get("latitude"), r.URL.Query().Get("longitude"))

	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&message.Text, &message.Latitude, &message.Longitude, &message.Expiry)
		messages = append(messages, message)
		if err != nil {
			log.Println(err.Error())
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(messages)
}

func main() {
	http.HandleFunc("/send/", sendHandler)
	http.HandleFunc("/messages/", messagesHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
