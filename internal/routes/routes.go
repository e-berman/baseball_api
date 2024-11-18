package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/e-berman/baseball_api/internal/db"
	"github.com/e-berman/baseball_api/internal/models"
)

// created for the purpose of referencing the server address and database connection
type Server struct {
	addr string
	db   db.DB
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
func NewServer(addr string, db db.DB) *Server {
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

	sm.HandleFunc("/api/position_players/", toHandleFunc(s.handlePositionPlayers))
	sm.HandleFunc("/api/pitchers/", toHandleFunc(s.handlePitchers))

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
func (s *Server) handlePositionPlayers(rw http.ResponseWriter, req *http.Request) error {
	if req.URL.Path == "/api/position_players/" {
		if req.Method == http.MethodGet {
			return s.handleGetPositionPlayers(rw, req)
		}
		if req.Method == http.MethodPost {
			return s.handleAddPositionPlayer(rw, req)
		}
	} else {
		if req.Method == http.MethodDelete {
			return s.handleDeletePositionPlayer(rw, req)
		}
		if req.Method == http.MethodGet {
			return s.handleGetPositionPlayerByID(rw, req)
		}
		if req.Method == http.MethodPut {
			return s.handleUpdatePositionPlayer(rw, req)
		}
	}

	return fmt.Errorf("invalid method for position players: %s", req.Method)
}

// handlePlayers handles the various routes given the respective request method
//
// conditionally separated based on whether a player id exists in the url or not
func (s *Server) handlePitchers(rw http.ResponseWriter, req *http.Request) error {
	if req.URL.Path == "/api/pitchers/" {
		if req.Method == http.MethodGet {
			return s.handleGetPitchers(rw, req)
		}
		if req.Method == http.MethodPost {
			return s.handleAddPitcher(rw, req)
		}
	} else {
		if req.Method == http.MethodDelete {
			return s.handleDeletePitcher(rw, req)
		}
		if req.Method == http.MethodGet {
			return s.handleGetPitcherByID(rw, req)
		}
		if req.Method == http.MethodPut {
			return s.handleUpdatePitcher(rw, req)
		}
	}

	return fmt.Errorf("invalid method for pitchers: %s", req.Method)
}

func (s *Server) handleGetPositionPlayers(rw http.ResponseWriter, req *http.Request) error {
	log.Println("GET all position players")
	players, err := s.db.GetPositionPlayers()
	if err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, players)
}

func (s *Server) handleGetPositionPlayerByID(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}

	player, err := s.db.GetPositionPlayerByID(id)
	if err != nil {
		return err
	}

	log.Println("GET player:", player.Name)

	return ToJSON(rw, http.StatusOK, player)
}

func (s *Server) handleAddPositionPlayer(rw http.ResponseWriter, req *http.Request) error {
	createPositionPlayerReq := models.CreatePositionPlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&createPositionPlayerReq); err != nil {
		return err
	}

	log.Println("POST player:", createPositionPlayerReq.Name)

	player := models.NewPositionPlayer(
		createPositionPlayerReq.Name,
		createPositionPlayerReq.Team,
		createPositionPlayerReq.G,
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

	if err := s.db.AddPositionPlayer(player); err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, createPositionPlayerReq)
}

func (s *Server) handleUpdatePositionPlayer(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}
	player, err := s.db.GetPositionPlayerByID(id)
	if err != nil {
		return err
	}

	updatePositionPlayerReq := models.UpdatePositionPlayerRequest{}
	if err := json.NewDecoder(req.Body).Decode(&updatePositionPlayerReq); err != nil {
		return err
	}

	player.Name = updatePositionPlayerReq.Name
	player.Team = updatePositionPlayerReq.Team
	player.G = updatePositionPlayerReq.G
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

	if err := s.db.UpdatePositionPlayer(player); err != nil {
		log.Println("error updating position player in database")
		return err
	}

	log.Println("UPDATE player id:", id)

	resMap := models.UpdatedPositionPlayer{
		UpdatedMap: map[string]int{
			"updated": id,
		},
	}

	return ToJSON(rw, http.StatusOK, resMap.UpdatedMap)
}

func (s *Server) handleDeletePositionPlayer(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}

	err = s.db.DeletePositionPlayer(id)
	if err != nil {
		return err
	}

	log.Println("DELETE player id:", id)

	resMap := models.DeletedPositionPlayer{
		DeletedMap: map[string]int{
			"deleted": id,
		},
	}
	return ToJSON(rw, http.StatusOK, resMap.DeletedMap)
}

func (s *Server) handleGetPitchers(rw http.ResponseWriter, req *http.Request) error {
	log.Println("GET all pitchers")
	players, err := s.db.GetPitchers()
	if err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, players)
}

func (s *Server) handleGetPitcherByID(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}

	player, err := s.db.GetPitcherByID(id)
	if err != nil {
		return err
	}

	log.Println("GET pitcher:", player.Name)

	return ToJSON(rw, http.StatusOK, player)
}

func (s *Server) handleAddPitcher(rw http.ResponseWriter, req *http.Request) error {
	createPitcherReq := models.CreatePitcherRequest{}
	if err := json.NewDecoder(req.Body).Decode(&createPitcherReq); err != nil {
		return err
	}

	log.Println("POST pitcher:", createPitcherReq.Name)

	player := models.NewPitcher(
		createPitcherReq.Name,
		createPitcherReq.Team,
		createPitcherReq.W,
		createPitcherReq.L,
		createPitcherReq.SV,
		createPitcherReq.G,
		createPitcherReq.GS,
		createPitcherReq.IP,
		createPitcherReq.K9,
		createPitcherReq.BB9,
		createPitcherReq.HR9,
		createPitcherReq.BABIP,
		createPitcherReq.LOB,
		createPitcherReq.GB,
		createPitcherReq.HRFB,
		createPitcherReq.VFA,
		createPitcherReq.ERA,
		createPitcherReq.XERA,
		createPitcherReq.FIP,
		createPitcherReq.XFIP,
		createPitcherReq.WAR,
	)

	if err := s.db.AddPitcher(player); err != nil {
		return err
	}

	return ToJSON(rw, http.StatusOK, createPitcherReq)
}

func (s *Server) handleUpdatePitcher(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}
	player, err := s.db.GetPitcherByID(id)
	if err != nil {
		return err
	}

	updatePitcherReq := models.UpdatePitcherRequest{}
	if err := json.NewDecoder(req.Body).Decode(&updatePitcherReq); err != nil {
		return err
	}

	player.Name = updatePitcherReq.Name
	player.Team = updatePitcherReq.Team
	player.W = updatePitcherReq.W
	player.L = updatePitcherReq.L
	player.SV = updatePitcherReq.SV
	player.G = updatePitcherReq.G
	player.GS = updatePitcherReq.GS
	player.IP = updatePitcherReq.IP
	player.K9 = updatePitcherReq.K9
	player.BB9 = updatePitcherReq.BB9
	player.HR9 = updatePitcherReq.HR9
	player.BABIP = updatePitcherReq.BABIP
	player.LOB = updatePitcherReq.LOB
	player.GB = updatePitcherReq.GB
	player.HRFB = updatePitcherReq.HRFB
	player.VFA = updatePitcherReq.VFA
	player.ERA = updatePitcherReq.ERA
	player.XERA = updatePitcherReq.XERA
	player.FIP = updatePitcherReq.FIP
	player.XFIP = updatePitcherReq.XFIP
	player.WAR = updatePitcherReq.WAR

	if err := s.db.UpdatePitcher(player); err != nil {
		return err
	}

	log.Println("UPDATE pitcher id:", id)

	resMap := models.UpdatedPitcher{
		UpdatedMap: map[string]int{
			"updated": id,
		},
	}

	return ToJSON(rw, http.StatusOK, resMap.UpdatedMap)
}

func (s *Server) handleDeletePitcher(rw http.ResponseWriter, req *http.Request) error {
	id, err := s.getIDFromPath(req)
	if err != nil {
		return err
	}

	err = s.db.DeletePitcher(id)
	if err != nil {
		return err
	}

	log.Println("DELETE pitcher id:", id)

	resMap := models.DeletedPitcher{
		DeletedMap: map[string]int{
			"deleted": id,
		},
	}
	return ToJSON(rw, http.StatusOK, resMap.DeletedMap)
}
