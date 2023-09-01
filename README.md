# baseball api

a REST API for baseball player stats written in Go.

uses the [net/http](https://pkg.go.dev/net/http) library for routing and the [pgx](https://github.com/jackc/pgx) library for the Postgres driver/toolkit.

the data in the database is structured based on the player data format according to [Fangraphs](https://www.fangraphs.com/)

the api and database are containerized via Docker.

OpenAPI specification available -> [HERE](https://app.swaggerhub.com/apis/e-berman/baseball-api/0.0.1)

## Dependencies

* Docker
* Docker compose

## To Run

1. Modify or add a desired csv to the baseball_api directory. It must be in the same format as the example stats.csv file. 

2. Create and build both the database and REST API containers: `make build`

3. If you do not want to import data via csv, skip steps 4-6. Use endpoints as desired.

4. Run `make get_path` to get the absolute filepath of desired csv file. 

5. Run the REST API container in an interactive shell and build the main executable: `make run`

6. Select 'Y' to add a filepath. When prompted, paste in the filepath from step 5.

7. When successfully imported, you can access the database with the following: `make db`


## Endpoints

| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | /api/position_players/ | retrieve all position players |
| GET | /api/position_players/:id | retrieve a single position player by id |
| POST | /api/position_players/ | add a new position player |
| PUT | /api/position_players/:id | edit field(s) of a single position player |
| DELETE | /api/position_players/:id | delete a single position player |
| --- | --- | --- |
| GET | /api/pitchers/ | retrieve all pitchers |
| GET | /api/pitchers/:id | retrieve a single pitcher by id |
| POST | /api/pitchers/ | add a new pitcher |
| PUT | /api/pitchers/:id | edit field(s) of a single pitcher |
| DELETE | /api/pitchers/:id | delete a single pitcher |


Example URL: `http://127.0.0.1:4242/api/position_players/`


Example JSON Payload for a position player:

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
    "weightedRunsCreatedPlus": 137,
    "walkRate": 7.7,
    "strikeoutRate": 4.2,
    "isolatedPower": 0.120,
    "battingAvgBallsInPlay": 0.294,
    "battingAvg": 0.338,
    "onBasePct": 0.388,
    "sluggingPct": 0.459,
    "weightedOnBaseAvg": 0.370,
    "expWeightedOnBaseAvg": 0.345,
    "baseRunning": 1.5,
    "winsAboveReplacement": 4.5,
}
```

## To Do 

- [x] add pitchers db table and api endpoints
- [ ] improve unit testing
- [ ] add daily csv import from reliable baseball statistics source (fangraphs or baseball savant?)
- [ ] build out front-end for documentation and api usage


