package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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
		p.l.Println("GET Players handler")
		p.getPlayers(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addPlayer(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		path_segments := strings.Split(r.URL.Path, "/")

		if len(path_segments) != 2 {
			p.l.Println("Invalid URI: no id field found")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		id_segment := path_segments[1]
		id, err := strconv.Atoi(id_segment)
		if err != nil {
			p.l.Println("Invalid URI: unable to convert string to int")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updatePlayer(id, rw, r)
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

func (p *Player) addPlayer(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Players handler")

	player := &data.Player{}
	err := player.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	
	data.AddPlayer(player)
}

func (p *Player) updatePlayer(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT Players handler")

	player := &data.Player{}
	err := player.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdatePlayer(id, player)
	if err == data.ErrPlayerNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}
