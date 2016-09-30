# Go Server for Message In A Bottle
## Usage
Send a message:
```
curl -i -H "Content-Type: application/json" \
-X POST \
-d '{"text":"xyz","latitude":119.123123,"longitude": 120.1222}' \
localhost:8080/send/
```

Get messages: `localhost:8080/messages/`

## Running
1. Postgres db: `pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start`
2. Go server: `go run server.go`

## Set Up
1. Install Go
2. Install postgres with homebrew and lib/pq (https://github.com/lib/pq)
3. Setup database `psql -f setup_db.sql`
