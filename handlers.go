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

	sm.HandleFunc("/players/", toHandleFunc(s.handlePlayers))
	
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
	if req.URL.Path == "/players/" {
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
	return nil
}

func (s *Server) handleAddPlayer(rw http.ResponseWriter, req *http.Request) error {
	createPlayerReq := CreatePlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&createPlayerReq); err != nil {
		return err
	}

	log.Println("POST player:", createPlayerReq.PlayerName)
	
	player := NewPlayer(
		createPlayerReq.PlayerName,
		createPlayerReq.Team,
		createPlayerReq.Position,
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
		log.Println(err)
		return err
	}

	return ToJSON(rw, http.StatusOK, player)
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
