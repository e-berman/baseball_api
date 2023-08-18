package main

// Player is the type used to represent a player
// swagger:model
type PositionPlayer struct {
	// Player id (auto incremented by database)
	ID		int     `json:"id"`
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	Games		int     `json:"games"`
	PA		int     `json:"plateAppearances"`
	HR		int     `json:"homeRuns"`
	R		int     `json:"runs"`
	RBI		int     `json:"runsBattedIn"`
	SB		int     `json:"stolenBases"`
	WRCPlus		int	`json:"weightedRunsCreatedPlus"`
	BbRate		float64 `json:"walkRate"`
	KRate		float64 `json:"strikeoutRate"`
	ISO		float64 `json:"isolatedPower"`
	AVG		float64 `json:"battingAvg"`
	OBP		float64 `json:"onBasePct"`
	SLG		float64 `json:"sluggingPct"`
	WOBA		float64 `json:"weightedOnBaseAvg"`
	XWOBA		float64 `json:"expWeightedOnBaseAvg"`
	BaseRunning	float64 `json:"baseRunning"`
	WAR		float64 `json:"winsAboveReplacement"`

}

// CreatePlayerRequest is the type used to create a player
// swagger:model
type CreatePositionPlayerRequest struct {
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	Games		int     `json:"games"`
	PA		int     `json:"plateAppearances"`
	HR		int     `json:"homeRuns"`
	R		int     `json:"runs"`
	RBI		int     `json:"runsBattedIn"`
	SB		int     `json:"stolenBases"`
	WRCPlus		int	`json:"weightedRunsCreatedPlus"`
	BbRate		float64 `json:"walkRate"`
	KRate		float64 `json:"strikeoutRate"`
	ISO		float64 `json:"isolatedPower"`
	AVG		float64 `json:"battingAvg"`
	OBP		float64 `json:"onBasePct"`
	SLG		float64 `json:"sluggingPct"`
	WOBA		float64 `json:"weightedOnBaseAvg"`
	XWOBA		float64 `json:"expWeightedOnBaseAvg"`
	BaseRunning	float64 `json:"baseRunning"`
	WAR		float64 `json:"winsAboveReplacement"`
}

// UpdatePlayerRequest is the type used to update a player
// swagger:model
type UpdatePositionPlayerRequest struct {
	Name		string  `json:"name"`
	Team		string  `json:"team"`
	Games		int     `json:"games"`
	PA		int     `json:"plateAppearances"`
	HR		int     `json:"homeRuns"`
	R		int     `json:"runs"`
	RBI		int     `json:"runsBattedIn"`
	SB		int     `json:"stolenBases"`
	WRCPlus		int	`json:"weightedRunsCreatedPlus"`
	BbRate		float64 `json:"walkRate"`
	KRate		float64 `json:"strikeoutRate"`
	ISO		float64 `json:"isolatedPower"`
	AVG		float64 `json:"battingAvg"`
	OBP		float64 `json:"onBasePct"`
	SLG		float64 `json:"sluggingPct"`
	WOBA		float64 `json:"weightedOnBaseAvg"`
	XWOBA		float64 `json:"expWeightedOnBaseAvg"`
	BaseRunning	float64 `json:"baseRunning"`
	WAR		float64 `json:"winsAboveReplacement"`
}

// UpdatedPlayer is the type used to return the id of the player updated upon success
// swagger:model
type UpdatedPositionPlayer struct {
	updatedMap map[string]int
}

// DeletedPlayer is the type used to return the id of the player deleted upon success
// swagger:model
type DeletedPositionPlayer struct {
	deletedMap map[string]int
}

func NewPositionPlayer(name, team string, games, pa, hr, r, rbi, sb, wrcPlus int, bbRate, kRate, iso, avg, obp, slg, woba, xWoba, bsr, war float64) *PositionPlayer {
	return &PositionPlayer{
		Name:		name,
		Team:		team,
		Games:		games,
		PA:		pa,
		HR:		hr,
		R:		r,
		RBI:		rbi,
		SB:		sb,
		WRCPlus:	wrcPlus,
		BbRate:		bbRate,
		KRate:		kRate,
		ISO:		iso,
		AVG:		avg,
		OBP:		obp,
		SLG:		slg,
		WOBA:		woba,
		XWOBA:		xWoba,
		BaseRunning:	bsr,
		WAR:		war,
	}
}
