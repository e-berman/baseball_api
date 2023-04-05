# baseball api

a REST API for baseball player stats.

written in Golang and uses the standard [net/http](https://pkg.go.dev/net/http) library for routing.

uses [pgx](https://github.com/jackc/pgx) as the Postgres driver/toolkit and two Docker containers for the api and database.

## Dependencies

* Docker
* Docker compose

## To Run

`docker compose up --build --wait`

URL: `http://127.0.0.1:4242/api/players/`

to view table data within container: `docker exec -it <container-name> psql -U <username> <database>`

### Endpoints

| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | /api/players/ | retrieve all players |
| GET | /api/causes/:id | retrieve a single player |
| POST | /api/players/ | add a new player |
| PUT | /api/players/:id | edit field(s) of a single player |
| DELETE | /api/players/:id | delete a single player |


