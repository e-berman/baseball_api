package data

import (
	"encoding/json"
	"time"
	"io"
	"fmt"
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

func (p *Player) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Players) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")
	return e.Encode(p)
}

func GetPlayers() Players {
	return playerList
}

func AddPlayer(p *Player) {
	p.ID = getNextID()
	playerList = append(playerList, p)
}

func UpdatePlayer(id int, p *Player) error {
	_, index, err := findPlayer(id)
	if err != nil {
		return err
	}

	p.ID = id
	playerList[index] = p

	return nil
}

var ErrPlayerNotFound = fmt.Errorf("Player not found")

func findPlayer(id int) (*Player, int, error) {
	for i, p := range playerList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrPlayerNotFound
}

func getNextID() int {
	last_player := playerList[len(playerList) - 1]
	return last_player.ID + 1
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
		ID:		2,
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
