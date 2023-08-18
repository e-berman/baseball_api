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
	// sm.HandleFunc("/api/players/import/", toHandleFunc(s.handleCSVImport))

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
			return s.handleAddPositionPlayer(rw, req)
		}
	} else {
		if req.Method == http.MethodDelete {
			return s.handleDeletePlayer(rw, req)
		}
		if req.Method == http.MethodGet {
			return s.handleGetPlayerByID(rw, req)
		}
		if req.Method == http.MethodPut {

			return s.handleUpdatePositionPlayer(rw, req)
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

func (s *Server) handleAddPositionPlayer(rw http.ResponseWriter, req *http.Request) error {
	createPositionPlayerReq := CreatePositionPlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&createPositionPlayerReq); err != nil {
		return err
	}

	log.Println("POST player:", createPositionPlayerReq.Name)

	player := NewPositionPlayer(
		createPositionPlayerReq.Name,
		createPositionPlayerReq.Team,
		createPositionPlayerReq.Games,
		createPositionPlayerReq.PA,
		createPositionPlayerReq.HR,
		createPositionPlayerReq.R,
		createPositionPlayerReq.RBI,
		createPositionPlayerReq.SB,
		createPositionPlayerReq.WRCPlus,
		createPositionPlayerReq.BbRate,
		createPositionPlayerReq.KRate,
		createPositionPlayerReq.ISO,
		createPositionPlayerReq.BABIP,
		createPositionPlayerReq.AVG,
		createPositionPlayerReq.OBP,
		createPositionPlayerReq.SLG,
		createPositionPlayerReq.WOBA,
		createPositionPlayerReq.XWOBA,
		createPositionPlayerReq.BsR,
		createPositionPlayerReq.WAR,
	)

	if err := s.db.AddPlayer(player); err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, createPositionPlayerReq)
}

func (s *Server) handleUpdatePositionPlayer(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}
	player, err := s.db.GetPlayerByID(id)
	if err != nil {
		return err
	}

	updatePositionPlayerReq := UpdatePositionPlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&updatePositionPlayerReq); err != nil {
		return err
	}

	player.Name = updatePositionPlayerReq.Name
	player.Team = updatePositionPlayerReq.Team
	player.Games = updatePositionPlayerReq.Games
	player.PA = updatePositionPlayerReq.PA
	player.HR = updatePositionPlayerReq.HR
	player.R = updatePositionPlayerReq.R
	player.RBI = updatePositionPlayerReq.RBI
	player.SB = updatePositionPlayerReq.SB
	player.WRCPlus = updatePositionPlayerReq.WRCPlus
	player.BbRate = updatePositionPlayerReq.BbRate
	player.KRate = updatePositionPlayerReq.KRate
	player.ISO = updatePositionPlayerReq.ISO
	player.BABIP = updatePositionPlayerReq.BABIP
	player.AVG = updatePositionPlayerReq.AVG
	player.OBP = updatePositionPlayerReq.OBP
	player.SLG = updatePositionPlayerReq.SLG
	player.WOBA = updatePositionPlayerReq.WOBA
	player.XWOBA = updatePositionPlayerReq.XWOBA
	player.BsR = updatePositionPlayerReq.BsR
	player.WAR = updatePositionPlayerReq.WAR

	if err := s.db.UpdatePlayer(player); err != nil {
		return err
	}

	log.Println("UPDATE player id:", id)

	resMap := UpdatedPositionPlayer{
		updatedMap: map[string]int{
			"updated": id,
		},
	}

	return ToJSON(rw, http.StatusOK, resMap.updatedMap)
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

	resMap := DeletedPositionPlayer{
		deletedMap: map[string]int{
			"deleted": id,
		},
	}
	return ToJSON(rw, http.StatusOK, resMap.deletedMap)
}
