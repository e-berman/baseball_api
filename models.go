package main

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

type Player struct {
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
