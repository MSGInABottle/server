# Go Server for Message In A Bottle
## Usage
Currently, we're running our Go server on an AWS EC2 instance at `52.41.253.190:9000`

Send a message:
```
curl -i -H "Content-Type: application/json" \
-X POST \
-d '{"text":"xyz","latitude":119.123123,"longitude": 120.1222}' \
52.41.253.190:9000/send/
```
Get messages: `curl 52.41.253.190:9000/messages/`

## Running Locally
1. Starting Postgres: `pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start`
2. Go server: `go run server.go`

## Set Up
1. Install Go
2. Install postgres with homebrew and lib/pq (https://github.com/lib/pq)
3. Setup database `psql -f setup_db.sql`
