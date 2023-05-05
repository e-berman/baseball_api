package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// created for the purpose of referencing the server address and database connection
type Server struct {
	addr string
	db   DB
}

// reduces code clutter for handleFunc
type apiFunc func(http.ResponseWriter, *http.Request) error

// error type for apiFunc
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

// ToJSON adds header and encodes to JSON
func ToJSON(rw http.ResponseWriter, status int, v any) error {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	return json.NewEncoder(rw).Encode(v)
}

// NewServer returns a Server struct given a passed server address and database connection
func NewServer(addr string, db DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

// StartServer starts the HTTP server and handles given routes
func (s *Server) StartServer() {
	sm := http.NewServeMux()
	server := &http.Server{
		Addr:         s.addr,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	sm.HandleFunc("/api/players/", toHandleFunc(s.handlePlayers))
	sm.HandleFunc("/api/players/import/", toHandleFunc(s.handleCSVImport))

	log.Println("Server started on port", server.Addr)

	log.Fatal(server.ListenAndServe())
}

// getIDFromPath returns a player id
//
// parses the player id from the url path
func (s *Server) getIDFromPath(req *http.Request) (int, error) {
	path_segments := strings.Split(req.URL.Path, "/")
	player_id_string := path_segments[len(path_segments)-1]
	player_id, err := strconv.Atoi(player_id_string)
	if err != nil {
		return -1, err
	}

	return player_id, nil
}

// handlePlayers handles the various routes given the respective request method
//
// conditionally separated based on whether a player id exists in the url or not
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

func (s *Server) handleCSVImport(rw http.ResponseWriter, req *http.Request) error {
	if req.URL.Path == "/api/players/import/" {
		if req.Method == http.MethodPost {
			return s.handleAddPlayersByCSVImport(rw, req)
		}
	}

	return fmt.Errorf("invalid method %s", req.Method)
}

func (s *Server) handleAddPlayersByCSVImport(rw http.ResponseWriter, req *http.Request) error {
	log.Println("POST players from CSV")
	players := ReadFromCSV()
	for _, player := range players {
		err := s.db.AddPlayer(player)
		if err != nil {
			return err
		}
	}

	return ToJSON(rw, http.StatusOK, players)
}

// swagger:route GET /api/players/ getPlayers
//
// # Returns all position players
//
// ---
// produces:
// - application/json
// responses:
//
//	200: Player
func (s *Server) handleGetPlayers(rw http.ResponseWriter, req *http.Request) error {
	log.Println("GET all players")
	players, err := s.db.GetPlayers()
	if err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, players)
}

// swagger:route GET /api/players/{id} getPlayerByID
//
// # Returns a Player by given id
//
// ---
// produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     schema:
//     	type: integer
//
// responses:
//
//	200: Player
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

// swagger:route POST /api/players/ addPlayer
//
// # Adds a Player to the database
//
// ---
// produces:
// - application/json
//
// responses:
//
//	200: CreatePlayerRequest
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
		createPlayerReq.BbRate,
		createPlayerReq.KRate,
		createPlayerReq.ISO,
		createPlayerReq.AVG,
		createPlayerReq.OBP,
		createPlayerReq.SLG,
		createPlayerReq.WOBA,
	)

	if err := s.db.AddPlayer(player); err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, createPlayerReq)
}

// swagger:route PUT /api/players/{id} updatePlayer
//
// # Updates a Player given an id
//
// ---
// produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     schema:
//     	type: integer
//
// responses:
//
//	200: UpdatePlayerRequest
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
	player.BbRate = updatePlayerReq.BbRate
	player.KRate = updatePlayerReq.KRate
	player.ISO = updatePlayerReq.ISO
	player.AVG = updatePlayerReq.AVG
	player.OBP = updatePlayerReq.OBP
	player.SLG = updatePlayerReq.SLG
	player.WOBA = updatePlayerReq.WOBA

	if err := s.db.UpdatePlayer(player); err != nil {
		return err
	}

	log.Println("UPDATE player id:", id)

	return ToJSON(rw, http.StatusOK, map[string]int{"updated": id})
}

// swagger:route DELETE /api/players/{id} deletePlayer
//
// # Deletes a Player from the database given id
//
// ---
// produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     schema:
//     	type: integer
//
// responses:
//
//	200: Player
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
