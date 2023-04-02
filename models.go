package main

import "time"

type CreatePlayerRequest struct {
	Name			string	`json:"name"`
	Team			string  `json:"team"`
	Games			int	`json:"games"`
	PA			int	`json:"plateAppearances"`
	HR			int	`json:"homeRuns"`
	R			int     `json:"runs"`
	RBI			int	`json:"runsBattedIn"`
	SB			int	`json:"stolenBases"`
	BbRate			float32	`json:"walkRate"`
	KRate			float32	`json:"strikeoutRate"`
	ISO			float32	`json:"isolatedPower"`
	BABIP			float32 `json:"battingAvgBallsInPlay"`
	AVG			float32	`json:"battingAvg"`
	OBP			float32	`json:"onBasePct"`
	SLG			float32	`json:"sluggingPct"`
	WOBA			float32	`json:"weightedOnBaseAvg"`
	WRCPlus			int	`json:"weightedRunsCreatedPlus"`	
	LastSeasonWAR		float32	`json:"lastSeasonWar"`
	CreatedOn		string	`json:"-"`
	UpdatedOn		string	`json:"-"`
	DeletedOn		string	`json:"-"`
}

type UpdatePlayerRequest struct {
	Name			string	`json:"name"`
	Team			string  `json:"team"`
	Games			int	`json:"games"`
	PA			int	`json:"plateAppearances"`
	HR			int	`json:"homeRuns"`
	R			int     `json:"runs"`
	RBI			int	`json:"runsBattedIn"`
	SB			int	`json:"stolenBases"`
	BbRate			float32	`json:"walkRate"`
	KRate			float32	`json:"strikeoutRate"`
	ISO			float32	`json:"isolatedPower"`
	BABIP			float32 `json:"battingAvgBallsInPlay"`
	AVG			float32	`json:"battingAvg"`
	OBP			float32	`json:"onBasePct"`
	SLG			float32	`json:"sluggingPct"`
	WOBA			float32	`json:"weightedOnBaseAvg"`
	WRCPlus			int	`json:"weightedRunsCreatedPlus"`	
	LastSeasonWAR		float32	`json:"lastSeasonWar"`
	CreatedOn		string	`json:"-"`
	UpdatedOn		string	`json:"-"`
	DeletedOn		string	`json:"-"`
}

type Player struct {
	ID			int	        `json:"id"`
	Name			string	        `json:"name"`
	Team			string          `json:"team"`
	Games			int	        `json:"games"`
	PA			int	        `json:"plateAppearances"`
	HR			int	        `json:"homeRuns"`
	R			int		`json:"runs"`
	RBI			int	        `json:"runsBattedIn"`
	SB			int	        `json:"stolenBases"`
	BbRate			float32	        `json:"walkRate"`
	KRate			float32	        `json:"strikeoutRate"`
	ISO			float32	        `json:"isolatedPower"`
	BABIP			float32         `json:"battingAvgBallsInPlay"`
	AVG			float32	        `json:"battingAvg"`
	OBP			float32	        `json:"onBasePct"`
	SLG			float32	        `json:"sluggingPct"`
	WOBA			float32	        `json:"weightedOnBaseAvg"`
	WRCPlus			int	        `json:"weightedRunsCreatedPlus"`	
	LastSeasonWAR		float32	        `json:"lastSeasonWar"`
	CreatedOn		time.Time	`json:"created_on"`
	UpdatedOn		time.Time	`json:"-"`
	DeletedOn		time.Time	`json:"-"`
}

func NewPlayer(name, team string, games, pa, hr, r, rbi, sb, wrcPlus int, bbRate, kRate, iso, babip, avg, obp, slg, woba, lastSeasonWAR float32) *Player {
	return &Player{
		Name: name,
		Team: team,
		Games: games,
		PA: pa,
		HR: hr,
		R: r,
		RBI: rbi,
		SB: sb,
		BbRate: bbRate,
		KRate: kRate,
		ISO: iso,
		BABIP: babip,
		AVG: avg,
		OBP: obp,
		SLG: slg,
		WOBA: woba,
		WRCPlus: wrcPlus,
		LastSeasonWAR: lastSeasonWAR,
	}
}
