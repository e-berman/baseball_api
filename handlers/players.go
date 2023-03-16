package handlers

import (
	"log"
	"net/http"
	"github.com/e-berman/baseball_api/data"
)
type Player struct {
	l *log.Logger
}

func NewPlayer(l *log.Logger) *Player {
	return &Player{l}
}

func (p *Player) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getPlayers(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Player) getPlayers(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetPlayers()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
