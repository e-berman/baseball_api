package main

// Player is the type used to represent a player
// swagger:model
type Player struct {
	// Player id (auto incremented by database)
	ID int `json:"id"`
	// Player name
	// required: true
	// min length: 3
	Name string `json:"name"`
	// Team of player
	// max length: 3
	// example: TBR
	Team string `json:"team"`
	// Number of games played in a season
	// min: 0
	// max: 162
	Games int `json:"games"`
	// Number of plate appearances in a season
	// min: 0
	// max: 800
	PA int `json:"plateAppearances"`
	// Number of home runs in a season
	// min: 0
	// max: 90
	HR int `json:"homeRuns"`
	// Number of runs in a season
	// min: 0
	// max: 200
	R int `json:"runs"`
	// Number of runs batted in in a season
	// min: 0
	// max: 200
	RBI int `json:"runsBattedIn"`
	// Number of stolen bases in a season
	// min: 0
	// max: 150
	SB int `json:"stolenBases"`
	// Rate at which a player walks in a season
	// min: 0
	// max: 100
	// example: 14.3
	BbRate float64 `json:"walkRate"`
	// Rate at which a player strikes out in a season
	// min: 0
	// max: 100
	// example: 20.7
	KRate float64 `json:"strikeoutRate"`
	// Raw power of a hitter ased on extra base hits and the type of extra base hit
	// min: 0
	// max: 1
	// example: 0.131
	ISO float64 `json:"isolatedPower"`
	// Batting average of a player in a season
	// min: 0
	// max: 1
	// example: 0.245
	AVG float64 `json:"battingAvg"`
	// Rate at which a player gets on base
	// min: 0
	// max: 1
	// example: 0.352
	OBP float64 `json:"onBasePct"`
	// Total number of bases a player records per at bat
	// min: 0
	// max: 1
	// example: 0.333
	SLG float64 `json:"sluggingPct"`
	// version of OBP that accounts for how the player got on base
	// min: 0
	// max: 1
	// example: 0.310
	WOBA float64 `json:"weightedOnBaseAvg"`
}

// CreatePlayerRequest is the type used to create a player
// swagger:model
type CreatePlayerRequest struct {
	// Player name
	// required: true
	// min length: 3
	Name string `json:"name"`
	// Team of player
	// max length: 3
	// example: TBR
	Team string `json:"team"`
	// Number of games played in a season
	// min: 0
	// max: 162
	Games int `json:"games"`
	// Number of plate appearances in a season
	// min: 0
	// max: 800
	PA int `json:"plateAppearances"`
	// Number of home runs in a season
	// min: 0
	// max: 90
	HR int `json:"homeRuns"`
	// Number of runs in a season
	// min: 0
	// max: 200
	R int `json:"runs"`
	// Number of runs batted in in a season
	// min: 0
	// max: 200
	RBI int `json:"runsBattedIn"`
	// Number of stolen bases in a season
	// min: 0
	// max: 150
	SB int `json:"stolenBases"`
	// Rate at which a player walks in a season
	// min: 0
	// max: 100
	// example: 14.3
	BbRate float64 `json:"walkRate"`
	// Rate at which a player strikes out in a season
	// min: 0
	// max: 100
	// example: 20.7
	KRate float64 `json:"strikeoutRate"`
	// Raw power of a hitter ased on extra base hits and the type of extra base hit
	// min: 0
	// max: 1
	// example: 0.131
	ISO float64 `json:"isolatedPower"`
	// Batting average of a player in a season
	// min: 0
	// max: 1
	// example: 0.245
	AVG float64 `json:"battingAvg"`
	// Rate at which a player gets on base
	// min: 0
	// max: 1
	// example: 0.352
	OBP float64 `json:"onBasePct"`
	// Total number of bases a player records per at bat
	// min: 0
	// max: 1
	// example: 0.333
	SLG float64 `json:"sluggingPct"`
	// version of OBP that accounts for how the player got on base
	// min: 0
	// max: 1
	// example: 0.310
	WOBA float64 `json:"weightedOnBaseAvg"`
}

// UpdatePlayerRequest is the type used to update a player
// swagger:model
type UpdatePlayerRequest struct {
	// Player name
	// required: true
	// min length: 3
	Name string `json:"name"`
	// Team of player
	// max length: 3
	// example: TBR
	Team string `json:"team"`
	// Number of games played in a season
	// min: 0
	// max: 162
	Games int `json:"games"`
	// Number of plate appearances in a season
	// min: 0
	// max: 800
	PA int `json:"plateAppearances"`
	// Number of home runs in a season
	// min: 0
	// max: 90
	HR int `json:"homeRuns"`
	// Number of runs in a season
	// min: 0
	// max: 200
	R int `json:"runs"`
	// Number of runs batted in in a season
	// min: 0
	// max: 200
	RBI int `json:"runsBattedIn"`
	// Number of stolen bases in a season
	// min: 0
	// max: 150
	SB int `json:"stolenBases"`
	// Rate at which a player walks in a season
	// min: 0
	// max: 100
	// example: 14.3
	BbRate float64 `json:"walkRate"`
	// Rate at which a player strikes out in a season
	// min: 0
	// max: 100
	// example: 20.7
	KRate float64 `json:"strikeoutRate"`
	// Raw power of a hitter ased on extra base hits and the type of extra base hit
	// min: 0
	// max: 1
	// example: 0.131
	ISO float64 `json:"isolatedPower"`
	// Batting average of a player in a season
	// min: 0
	// max: 1
	// example: 0.245
	AVG float64 `json:"battingAvg"`
	// Rate at which a player gets on base
	// min: 0
	// max: 1
	// example: 0.352
	OBP float64 `json:"onBasePct"`
	// Total number of bases a player records per at bat
	// min: 0
	// max: 1
	// example: 0.333
	SLG float64 `json:"sluggingPct"`
	// version of OBP that accounts for how the player got on base
	// min: 0
	// max: 1
	// example: 0.310
	WOBA float64 `json:"weightedOnBaseAvg"`
}

// UpdatedPlayer is the type used to return the id of the player updated upon success
// swagger:model
type UpdatedPlayer struct {
	// Map of "updated" key and id value
	// example: [{"updated": 1}]
	updatedMap map[string]int
}

// AddedPlayers is the type used to return the number of players imported via .csv
// swagger:model
type AddedPlayers struct {
	// Map of "added" key and number of players added value
	// example: [{"added": 35}]
	addedMap map[string]int
}

// DeletedPlayer is the type used to return the id of the player deleted upon success
// swagger:model
type DeletedPlayer struct {
	// Map of "deleted" key and id of player deleted value
	// example: [{"deleted": 3}]
	deletedMap map[string]int
}

func NewPlayer(name, team string, games, pa, hr, r, rbi, sb int, bbRate, kRate, iso, avg, obp, slg, woba float64) *Player {
	return &Player{
		Name:   name,
		Team:   team,
		Games:  games,
		PA:     pa,
		HR:     hr,
		R:      r,
		RBI:    rbi,
		SB:     sb,
		BbRate: bbRate,
		KRate:  kRate,
		ISO:    iso,
		AVG:    avg,
		OBP:    obp,
		SLG:    slg,
		WOBA:   woba,
	}
}
