# baseball api

a REST API for baseball player stats written in Go.

uses the [net/http](https://pkg.go.dev/net/http) library for routing and the [pgx](https://github.com/jackc/pgx) library for the Postgres driver/toolkit.

the data in the database is structured based on the player data format according to [Fangraphs](https://www.fangraphs.com/)

the api and database are containerized via Docker.

OpenAPI specification available -> [HERE](https://app.swaggerhub.com/apis/e-berman/baseball-api/0.0.1)

## Dependencies

* docker
* docker-compose

## To Run

1. Modify or add a desired csv to the `baseball_api/assets` directory. It must be in the same format as the batters.csv and pitchers.csv file(s). 

2. Create and build both the database and REST API containers: `make build`

3. Database will import .csv data if added. You can access the database with the following: `make db`


Example URL: `http://localhost:4242/api/position_players/`

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
Example JSON Payload for a pitcher:
```
{
    "id": 1,
    "name": "Aaron Nola",
    "team": "PHI",
    "wins": 11,
    "losses": 13,
    "saves": 0,
    "games": 32,
    "gamesSaved": 32,
    "inningsPitched": 205,
    "strikeoutsPerNine": 10.32,
    "walksPerNine": 1.27,
    "homeRunsPerNine": 0.83,
    "battingAvgBallsInPlay": 0.289,
    "leftOnBase": 73,
    "groundballRate": 43.6,
    "homeRunToFlyBallRatio": 9.8,
    "fourseamFastballVelocity": 92.9,
    "earnedRunAvg": 3.25,
    "expectedEarnedRunAvg": 2.74,
    "fielderIndependentPitching": 2.58,
    "expectedFielderIndependentPitching": 2.77,
    "winsAboveReplacement": 6.3
}
```

## To Do 

- [x] add pitchers db table and api endpoints
- [x] refactor Dockerfile to multi-stage build
- [ ] improve unit testing
- [ ] add daily csv import from reliable baseball statistics source (fangraphs or baseball savant?)
- [ ] build out front-end for documentation and api usage


