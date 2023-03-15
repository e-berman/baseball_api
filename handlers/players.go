package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)
type Player struct {
	l *log.Logger
}

func NewPlayer(l *log.Logger) *Player {
	return &Player{l}
}

func (p *Player) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
