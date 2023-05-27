package main

// Player is the type used to represent a player
// swagger:model
type Player struct {
	// Player id (auto incremented by database)
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Team   string  `json:"team"`
	Games  int     `json:"games"`
	PA     int     `json:"plateAppearances"`
	HR     int     `json:"homeRuns"`
	R      int     `json:"runs"`
	RBI    int     `json:"runsBattedIn"`
	SB     int     `json:"stolenBases"`
	BbRate float64 `json:"walkRate"`
	KRate  float64 `json:"strikeoutRate"`
	ISO    float64 `json:"isolatedPower"`
	AVG    float64 `json:"battingAvg"`
	OBP    float64 `json:"onBasePct"`
	SLG    float64 `json:"sluggingPct"`
	WOBA   float64 `json:"weightedOnBaseAvg"`
}

// CreatePlayerRequest is the type used to create a player
// swagger:model
type CreatePlayerRequest struct {
	Name   string  `json:"name"`
	Team   string  `json:"team"`
	Games  int     `json:"games"`
	PA     int     `json:"plateAppearances"`
	HR     int     `json:"homeRuns"`
	R      int     `json:"runs"`
	RBI    int     `json:"runsBattedIn"`
	SB     int     `json:"stolenBases"`
	BbRate float64 `json:"walkRate"`
	KRate  float64 `json:"strikeoutRate"`
	ISO    float64 `json:"isolatedPower"`
	AVG    float64 `json:"battingAvg"`
	OBP    float64 `json:"onBasePct"`
	SLG    float64 `json:"sluggingPct"`
	WOBA   float64 `json:"weightedOnBaseAvg"`
}

// UpdatePlayerRequest is the type used to update a player
// swagger:model
type UpdatePlayerRequest struct {
	Name   string  `json:"name"`
	Team   string  `json:"team"`
	Games  int     `json:"games"`
	PA     int     `json:"plateAppearances"`
	HR     int     `json:"homeRuns"`
	R      int     `json:"runs"`
	RBI    int     `json:"runsBattedIn"`
	SB     int     `json:"stolenBases"`
	BbRate float64 `json:"walkRate"`
	KRate  float64 `json:"strikeoutRate"`
	ISO    float64 `json:"isolatedPower"`
	AVG    float64 `json:"battingAvg"`
	OBP    float64 `json:"onBasePct"`
	SLG    float64 `json:"sluggingPct"`
	WOBA   float64 `json:"weightedOnBaseAvg"`
}

// UpdatedPlayer is the type used to return the id of the player updated upon success
// swagger:model
type UpdatedPlayer struct {
	updatedMap map[string]int
}

// DeletedPlayer is the type used to return the id of the player deleted upon success
// swagger:model
type DeletedPlayer struct {
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
