# Go Server for Message In A Bottle

## Running
1. Postgres db: `pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start`
2. Go server: `go run server.go`

## Set Up
1. Install Go
2. Install postgres with homebrew and lib/pq (https://github.com/lib/pq)
3. Setup database `psql -f setup_db.sql`
