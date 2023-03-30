package main

import "time"

type CreatePlayerRequest struct {
	PlayerName		string	`json:"playerName"`
	Team			string  `json:"team"`
	Position		string	`json:"position"`
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
	PlayerName		string	`json:"playerName"`
	Team			string  `json:"team"`
	Position		string	`json:"position"`
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
	PlayerName		string	        `json:"playerName"`
	Team			string          `json:"team"`
	Position		string	        `json:"position"`
	Games			int	        `json:"games"`
	PA			int	        `json:"plateAppearances"`
	HR			int	        `json:"homeRuns"`
	R			int             `json:"runs"`
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

type Team struct {
	ID			int	        `json:"id"`
	Team			string          `json:"team"`
	Games			int	        `json:"games"`
	PA			int	        `json:"plate_appearances"`
	HR			int	        `json:"home_runs"`
	R			int             `json:"runs"`
	RBI			int	        `json:"runs_batted_in"`
	SB			int	        `json:"stolen_bases"`
	BbRate			float32	        `json:"walk_rate(BB%)"`
	KRate			float32	        `json:"strikeout_rate(K%)"`
	ISO			float32	        `json:"isolated_power"`
	BABIP			float32         `json:"batting_avg_balls_in_play"`
	AVG			float32	        `json:"batting_avg"`
	OBP			float32	        `json:"on_base_pct"`
	SLG			float32	        `json:"slugging_pct"`
	WOBA			float32	        `json:"weighted_on_base_avg"`
	WRCPlus			int	        `json:"weighted_runs_created_plus"`	
	TeamWAR			float32	        `json:"team_war"`
	CreatedOn		time.Time	`json:""`
	UpdatedOn		time.Time	`json:"-"`
	DeletedOn		time.Time	`json:"-"`
}

func NewPlayer(playerName, team, position string, games, pa, hr, r, rbi, sb, wrcPlus int, bbRate, kRate, iso, babip, avg, obp, slg, woba, lastSeasonWAR float32) *Player {
	return &Player{
		PlayerName: playerName,
		Team:	team,
		Position: position,
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
		CreatedOn: time.Now().UTC(),
	}
}
