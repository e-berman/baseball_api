package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"time"
	"strconv"
	"strings"
)

type Server struct {
	addr	 string
	db       DB
}

// references the Player model in models.go
type PlayerModel struct {
	player *Player
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiErr struct {
	Error string
}

// decorates apiFunc and handles error to reduce code clutter.
// returns an http.HandlerFunc
func toHandleFunc(f apiFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if err := f(rw, req); err != nil {
			ToJSON(rw, http.StatusBadRequest, apiErr{Error: err.Error()})
		}
	}
}

// encodes to JSON and writes/sets header
func ToJSON(rw http.ResponseWriter, status int, v any) error {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	return json.NewEncoder(rw).Encode(v)
}

func NewServer(addr string, db DB) *Server {
	return &Server{
		addr: addr,
		db: db,
	}
}

// starts server and lists routes
func (s *Server) StartServer() {
	sm := http.NewServeMux()
	server := &http.Server{
		Addr: s.addr,
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	sm.HandleFunc("/api/players/", toHandleFunc(s.handlePlayers))
	
	log.Println("Server started on port", server.Addr)

	log.Fatal(server.ListenAndServe())
}

func (s *Server) getIDFromPath(req *http.Request) (int, error) {
	path_segments := strings.Split(req.URL.Path, "/")
	player_id_string := path_segments[len(path_segments)-1]
	player_id, err := strconv.Atoi(player_id_string)
	if err != nil {
		return -1, err
	}

	return player_id, nil
}

func (s *Server) handlePlayers(rw http.ResponseWriter, req *http.Request) error {
	if req.URL.Path == "/api/players/" {
		if req.Method == http.MethodGet {
			return s.handleGetPlayers(rw, req)
		}
		if req.Method == http.MethodPost {
			return s.handleAddPlayer(rw, req)
		}
	} else {
		if req.Method == http.MethodDelete {
			return s.handleDeletePlayer(rw, req)
		}
		if req.Method == http.MethodGet {
			return s.handleGetPlayerByID(rw, req)
		}
		if req.Method == http.MethodPut {
			return s.handleUpdatePlayer(rw, req)
		}
	}	
	return fmt.Errorf("invalid method %s", req.Method)
}

func (s *Server) handleGetPlayers(rw http.ResponseWriter, req *http.Request) error {
	log.Println("GET all players")
	players, err := s.db.GetPlayers()
	if err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, players)
}

func (s *Server) handleGetPlayerByID(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}
	
	player, err := s.db.GetPlayerByID(id)
	if err != nil {
		return err
	}

	log.Println("GET player:", player.Name)

	return ToJSON(rw, http.StatusOK, player)	
}

func (s *Server) handleAddPlayer(rw http.ResponseWriter, req *http.Request) error {
	createPlayerReq := CreatePlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&createPlayerReq); err != nil {
		return err
	}

	log.Println("POST player:", createPlayerReq.Name)
	
	player := NewPlayer(
		createPlayerReq.Name,
		createPlayerReq.Team,
		createPlayerReq.Games,
		createPlayerReq.PA,
		createPlayerReq.HR,
		createPlayerReq.R,
		createPlayerReq.RBI,
		createPlayerReq.SB,
		createPlayerReq.WRCPlus,
		createPlayerReq.BbRate,
		createPlayerReq.KRate,
		createPlayerReq.ISO,
		createPlayerReq.BABIP,
		createPlayerReq.AVG,
		createPlayerReq.OBP,
		createPlayerReq.SLG,
		createPlayerReq.WOBA,
		createPlayerReq.LastSeasonWAR,
	)

	if err := s.db.AddPlayer(player); err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, createPlayerReq)
}

func (s *Server) handleUpdatePlayer(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}
	player, err := s.db.GetPlayerByID(id)
	if err != nil {
		return err
	}

	updatePlayerReq := UpdatePlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&updatePlayerReq); err != nil {
		return err
	}

	player.Name = updatePlayerReq.Name
	player.Team = updatePlayerReq.Team
	player.Games = updatePlayerReq.Games
	player.PA = updatePlayerReq.PA
	player.HR = updatePlayerReq.HR
	player.R = updatePlayerReq.R
	player.RBI = updatePlayerReq.RBI
	player.SB = updatePlayerReq.SB
	player.WRCPlus = updatePlayerReq.WRCPlus
	player.BbRate = updatePlayerReq.BbRate
	player.KRate = updatePlayerReq.KRate
	player.ISO = updatePlayerReq.ISO
	player.BABIP = updatePlayerReq.BABIP
	player.AVG = updatePlayerReq.AVG
	player.OBP = updatePlayerReq.OBP
	player.SLG = updatePlayerReq.SLG
	player.WOBA = updatePlayerReq.WOBA
	player.LastSeasonWAR = updatePlayerReq.LastSeasonWAR

	if err := s.db.UpdatePlayer(player); err != nil {
		return err
	}

	log.Println("UPDATE player id:", id)

	return ToJSON(rw, http.StatusOK, map[string]int{"updated": id})
}

func (s *Server) handleDeletePlayer(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}

	err = s.db.DeletePlayer(id)
	if err != nil {
		return err
	}

	log.Println("DELETE player id:", id)

	return ToJSON(rw, http.StatusOK, map[string]int{"deleted": id})
}
