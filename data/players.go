package data

import (
	"encoding/json"
	"time"
	"io"
)

type Player struct {
	ID		int	`json:"id"`
	Name		string	`json:"name"`
	TeamAbbrev	string  `json:"team_abbrev"`
	Position	string	`json:"position"`
	LastSeasonWAR	float32	`json:"last_season_war"`
	BattingAvg	float32	`json:"batting_avg"`
	HomeRuns	int8	`json:"home_runs"`
	StolenBases	int8	`json:"stolen_bases"`
	CreatedOn	string	`json:"-"`
	UpdatedOn	string	`json:"-"`
	DeletedOn	string	`json:"-"`
}

type Players []*Player

func (p *Players) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")
	return e.Encode(p)
}

func GetPlayers() Players {
	return playerList
}

var playerList = []*Player{
	{
		ID:		1,
		Name:		"Juan Soto",
		TeamAbbrev:	"SDP",
		Position:	"RF/LF",
		LastSeasonWAR:  3.8,
		BattingAvg:	.242,
		HomeRuns:	27,
		StolenBases:	6,
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:	time.Now().UTC().String(),
	},
	{
		ID:		1,
		Name:		"Mike Trout",
		TeamAbbrev:	"LAA",
		Position:	"CF",
		LastSeasonWAR:  6.0,
		BattingAvg:	.283,
		HomeRuns:	40,
		StolenBases:	1,
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:	time.Now().UTC().String(),
	},
}
