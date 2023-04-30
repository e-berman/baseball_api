# baseball api

a REST API for baseball player stats.

written in Golang and uses the standard [net/http](https://pkg.go.dev/net/http) library for routing.

uses [pgx](https://github.com/jackc/pgx) as the Postgres driver/toolkit and two Docker containers for the api and database.

## Dependencies

* Docker
* Docker compose

## To Run

1. Modify or add a desired csv to the baseball_api directory. It must be in the same format as the example stats.csv file. 

2. Create and build both the database and REST API containers:<br>
`docker compose up --build`

3. If you do not want to import data via csv, skip steps 4-6. Use endpoints as desired.

4. Run the REST API container in an interactive shell and build the main executable:<br>
`docker compose exec rest_api sh`

5. Run `echo $(realpath *filename*)` to get the absolute filepath of desired csv file. 

6. Within the container, start the executable with:<br>
`./main`

7. Select 'Y' to add a filepath. When prompted, paste in the filepath from step 5.

8. When successfully imported, you can access the database with the following:<br>
`docker exec -it db psql -U <username> <database>`


## Endpoints

| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | /api/players/ | retrieve all players |
| GET | /api/causes/:id | retrieve a single player |
| POST | /api/players/ | add a new player |
| POST | /api/players/external/ | import a list of players via csv |
| PUT | /api/players/:id | edit field(s) of a single player |
| DELETE | /api/players/:id | delete a single player |

Example URL: `http://127.0.0.1:4242/api/players/`

Example JSON Payload:

```
{
    "name": "Tony Gwynn",
    "team": "SDP",
    "games": 2440,
    "plateAppearances": 10232,
    "homeRuns": 135,
    "runs": 1383,
    "runsBattedIn": 1138,
    "stolenBases": 319,
    "walkRate": 7.7,
    "strikeoutRate": 4.2,
    "isolatedPower": 0.120,
    "battingAvg": 0.338,
    "onBasePct": 0.388,
    "sluggingPct": 0.459,
    "weightedOnBaseAvg": 0.370,
}
