package db

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"math"
	
	"github.com/e-berman/baseball_api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// creates the DB type for handlers.go to utilize in the various routes
type DB interface {
	AddPositionPlayer(*models.PositionPlayer) error
	DeletePositionPlayer(int) error
	UpdatePositionPlayer(*models.PositionPlayer) error
	GetPositionPlayers() ([]*models.PositionPlayer, error)
	GetPositionPlayerByID(int) (*models.PositionPlayer, error)
	AddPitcher(*models.Pitcher) error
	DeletePitcher(int) error
	UpdatePitcher(*models.Pitcher) error
	GetPitchers() ([]*models.Pitcher, error)
	GetPitcherByID(int) (*models.Pitcher, error)
}

// Holds the pgxpool.Pool type for the initialization of the Postgres database via the pgx driver
type DBPool struct {
	Poolconn *pgxpool.Pool
}

// NewDBPool returns a DBPool instance
//
// sets up a new Postgres pgx pool
// if pgx pool is unable to be set up, will fail
func NewDBPool() (*DBPool, error) {
	godotenv.Load()
	db_url := os.Getenv("POSTGRES_URL")

	pool, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}

	log.Println("Connected to db on port:", os.Getenv("PORT"))

	return &DBPool{
		Poolconn: pool,
	}, nil
}

// InitializePlayerTable will create the position_players postgres table upon initialization
func (pool *DBPool) InitializePositionPlayerTable() error {
	return pool.CreatePositionPlayerTable()
}

func (pool *DBPool) InitializePitcherTable() error {
	return pool.CreatePitcherTable()
}

// CreatePlayerTable executes the query to create the position_player table with pgx
func (pool *DBPool) CreatePositionPlayerTable() error {
	query := `create table if not exists position_players (
		player_id serial primary key NOT NULL,
		name text NOT NULL,
		team text,
		g int CHECK (g >= 0),
		pa int CHECK (pa >= 0),
		hr int CHECK (hr >= 0),
		runs int CHECK (runs >= 0),
		rbi int CHECK (rbi >= 0),
		sb int CHECK (sb >= 0),
		wrc_plus int CHECK (wrc_plus >= 0),
		bb_rate float8 CHECK (bb_rate >= 0),
		k_rate float8 CHECK (k_rate >= 0),
		iso float8 CHECK (iso >= 0),
	        babip float8 CHECK (babip >= 0),
		average float8 CHECK (average >= 0),
		obp float8 CHECK (obp >= 0),
		slg float8 CHECK (slg >= 0),
		woba float8 CHECK (woba >= 0),
		x_woba float8 CHECK (x_woba >= 0),
		bsr float8,
		war float8,
		unique (name, team))`

	_, err := pool.Poolconn.Exec(context.Background(), query)

	return err
}

func (pool *DBPool) CreatePitcherTable() error {
	query := `CREATE table IF NOT EXISTS pitchers (
		player_id serial primary key NOT NULL,
		name text NOT NULL,
		team text,
		w int CHECK (w >= 0),
		l int CHECK (l >= 0),
	        sv int CHECK (sv >= 0),
		g int CHECK (g >= 0),
		gs int CHECK (gs >= 0),
		ip float8 CHECK (ip >= 0),
		k9 float8 CHECK (k9 >= 0),
		bb9 float8 CHECK (bb9 >= 0),
		hr9 float8 CHECK (hr9 >= 0),
		babip float8 CHECK (babip >= 0),
		lob float8 CHECK (lob >= 0),
		gb float8 CHECK (gb >= 0),
		hrfb float8 CHECK (hrfb >= 0),
		vfa float8 CHECK (vfa >= 0),
		era float8 CHECK (era >= 0),
		xera float8 CHECK (xera >= 0),
		fip float8 CHECK (fip >= 0),
		xfip float8 CHECK (xfip >= 0),
		war float8,
		unique (name, team))`

	_, err := pool.Poolconn.Exec(context.Background(), query)

	return err
}

// ClearPlayerTable clears the position_players table for testing purposes
func (pool *DBPool) ClearPlayerTable() error {
	clear_records_query := `DROP * FROM position_players`
	reset_id_query := `ALTER SEQUENCE id RESTART WITH 1`

	_, err_clear := pool.Poolconn.Exec(context.Background(), clear_records_query)
	if err_clear != nil {
		return err_clear
	}
	_, err_reset := pool.Poolconn.Exec(context.Background(), reset_id_query)
	if err_reset != nil {
		return err_reset
	}

	return nil
}

func ConvertFloatToInt(record string) int {
	val := ConvertToFloat(record)

	return int(val)
}

func ConvertToInt(record string) int {
	val, err := strconv.Atoi(record)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func ConvertToFloat(record string) float64 {
	val, err := strconv.ParseFloat(record, 64)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

// round function used from:
// https://gosamples.dev/round-float/
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ReadFromCSVPositionPlayer() []*models.PositionPlayer {
	file, err := os.Open("./assets/batters.csv")
	if err != nil {
		log.Fatalf("Unable to open CSV file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read CSV file: %v\n", err)
	}

	var players []*models.PositionPlayer

	for i, record := range records {
		if i == 0 {
			continue
		}

		adjusted_bb_rate := ConvertToFloat(record[9]) * 100
		adjusted_k_rate := ConvertToFloat(record[10])* 100

		row := &models.PositionPlayer{
			Name:		record[0],
			Team:		record[1],
			G:		ConvertToInt(record[2]),
			PA:		ConvertToInt(record[3]),
			HR:		ConvertToInt(record[4]),
			R:		ConvertToInt(record[5]),
			RBI:		ConvertToInt(record[6]),
			SB:		ConvertToInt(record[7]),
			WRCPlus:	ConvertFloatToInt(record[8]),
			BbRate:		roundFloat(adjusted_bb_rate, 1),
			KRate:		roundFloat(adjusted_k_rate, 1),
			ISO:		roundFloat(ConvertToFloat(record[11]), 3),
			BABIP:		roundFloat(ConvertToFloat(record[12]), 3),
			AVG:		roundFloat(ConvertToFloat(record[13]), 3),
			OBP:		roundFloat(ConvertToFloat(record[14]), 3),
			SLG:		roundFloat(ConvertToFloat(record[15]), 3),
			WOBA:		roundFloat(ConvertToFloat(record[16]), 3),
			XWOBA:		roundFloat(ConvertToFloat(record[17]), 3),
			BsR:		roundFloat(ConvertToFloat(record[18]), 1),
			WAR:		roundFloat(ConvertToFloat(record[19]), 1),
		}
		players = append(players, row)
	}

	return players
}

func ReadFromCSVPitcher() []*models.Pitcher {
	file, err := os.Open("../../assets/pitchers.csv")
	if err != nil {
		log.Fatalf("Unable to open CSV file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read CSV file: %v\n", err)
	}

	var players []*models.Pitcher

	for i, record := range records {
		if i == 0 {
			continue
		}

		adjusted_lob := ConvertToFloat(record[12]) * 100
		adjusted_gb_rate := ConvertToFloat(record[13]) * 100
		adjusted_hrfb_rate := ConvertToFloat(record[14]) * 100

		row := &models.Pitcher{
			Name:		record[0],
			Team:		record[1],
			W:		ConvertToInt(record[2]),
			L:		ConvertToInt(record[3]),
			SV:		ConvertToInt(record[4]),
			G:		ConvertToInt(record[5]),
			GS:		ConvertToInt(record[6]),
			IP:		ConvertToFloat(record[7]),
			K9:		roundFloat(ConvertToFloat(record[8]), 2),
			BB9:		roundFloat(ConvertToFloat(record[9]), 2),
			HR9:		roundFloat(ConvertToFloat(record[10]), 2),
			BABIP:		roundFloat(ConvertToFloat(record[11]), 3),
			LOB:		roundFloat(adjusted_lob, 1),
			GB:		roundFloat(adjusted_gb_rate, 1),
			HRFB:		roundFloat(adjusted_hrfb_rate, 1),
			VFA:		roundFloat(ConvertToFloat(record[15]), 1),
			ERA:		roundFloat(ConvertToFloat(record[16]), 2),
			XERA:		roundFloat(ConvertToFloat(record[17]), 2),
			FIP:		roundFloat(ConvertToFloat(record[18]), 2),
			XFIP:		roundFloat(ConvertToFloat(record[19]), 2),
			WAR:		roundFloat(ConvertToFloat(record[20]), 1),
		}
		players = append(players, row)
	}

	return players
}

func (pool *DBPool) ImportPositionPlayerDataFromCSV() error {
	players := ReadFromCSVPositionPlayer()
	for _, player := range players {
		err := pool.AddPositionPlayer(player)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pool *DBPool) ImportPitcherDataFromCSV() error {
	players := ReadFromCSVPitcher()
	for _, player := range players {
		err := pool.AddPitcher(player)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetPlayers will return a list of players
//
// retrieves all existing players in the position_players table
func (pool *DBPool) GetPositionPlayers() ([]*models.PositionPlayer, error) {
	query := `SELECT * FROM position_players`

	rows, err := pool.Poolconn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	players := []*models.PositionPlayer{}
	for rows.Next() {
		player := &models.PositionPlayer{}
		err := rows.Scan(
			&player.ID,
			&player.Name,
			&player.Team,
			&player.G,
			&player.PA,
			&player.HR,
			&player.R,
			&player.RBI,
			&player.SB,
			&player.WRCPlus,
			&player.BbRate,
			&player.KRate,
			&player.ISO,
			&player.BABIP,
			&player.AVG,
			&player.OBP,
			&player.SLG,
			&player.WOBA,
			&player.XWOBA,
			&player.BsR,
			&player.WAR,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

// GetPlayerByID will return a player
//
// it will query the position_players table based on a given player id
func (pool *DBPool) GetPositionPlayerByID(id int) (*models.PositionPlayer, error) {
	query := `SELECT * FROM position_players WHERE player_id = $1`
	player := &models.PositionPlayer{}

	err := pool.Poolconn.QueryRow(context.Background(), query, id).Scan(
		&player.ID,
		&player.Name,
		&player.Team,
		&player.G,
		&player.PA,
		&player.HR,
		&player.R,
		&player.RBI,
		&player.SB,
		&player.WRCPlus,
		&player.BbRate,
		&player.KRate,
		&player.ISO,
		&player.BABIP,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
		&player.XWOBA,
		&player.BsR,
		&player.WAR,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return player, nil
}

// AddPlayer will add a player to the position_players table
//
// will return nil if successful, error if unsuccessful
func (pool *DBPool) AddPositionPlayer(player *models.PositionPlayer) error {
	query := `INSERT INTO position_players 
	(name, team, g, pa, hr, runs, rbi, sb, wrc_plus, bb_rate, k_rate, iso, babip, average, obp, slg, woba, x_woba, bsr, war)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
	ON CONFLICT (name, team) DO NOTHING`

	_, err := pool.Poolconn.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.G,
		&player.PA,
		&player.HR,
		&player.R,
		&player.RBI,
		&player.SB,
		&player.WRCPlus,
		&player.BbRate,
		&player.KRate,
		&player.ISO,
		&player.BABIP,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
		&player.XWOBA,
		&player.BsR,
		&player.WAR,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePlayer will update a player in the position_players table given a player id
//
// if unsuccessful, will return an error
func (pool *DBPool) UpdatePositionPlayer(player *models.PositionPlayer) error {
	query := `UPDATE position_players SET
	name = $1,
	team = $2,
	g = $3,
	pa = $4,
	hr = $5,
	runs = $6,
	rbi = $7,
	sb = $8,
	wrc_plus = $9,
	bb_rate = $10,
	k_rate = $11,
	iso = $12,
	babip = $13,
	average = $13,
	obp = $14,
	slg = $15,
	woba = $16,
	x_woba = $17,
	bsr = $18,
	war = $19,
	WHERE player_id = $20`

	res, err := pool.Poolconn.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.G,
		&player.PA,
		&player.HR,
		&player.R,
		&player.RBI,
		&player.SB,
		&player.WRCPlus,
		&player.BbRate,
		&player.KRate,
		&player.ISO,
		&player.BABIP,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
		&player.XWOBA,
		&player.BsR,
		&player.WAR,
		&player.ID,
	)
	if err != nil {
		return err
	}

	log.Println("rows affected:", res.RowsAffected())

	return nil
}

// DeletePlayer deletes a player by player id in the position_players table
func (pool *DBPool) DeletePositionPlayer(id int) error {
	query := `DELETE FROM position_players WHERE player_id = $1`

	_, err := pool.Poolconn.Exec(context.Background(), query, id)
	return err
}


// *******************
// Pitcher methods
// *******************

// GetPitchers will return a list of pitchers
//
// retrieves all existing players in the pitchers table
func (pool *DBPool) GetPitchers() ([]*models.Pitcher, error) {
	query := `SELECT * FROM pitchers`

	rows, err := pool.Poolconn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	players := []*models.Pitcher{}
	for rows.Next() {
		player := &models.Pitcher{}
		err := rows.Scan(
			&player.ID,
			&player.Name,
			&player.Team,
			&player.W,
			&player.L,
			&player.SV,
			&player.G,
			&player.GS,
			&player.IP,
			&player.K9,
			&player.BB9,
			&player.HR9,
			&player.BABIP,
			&player.LOB,
			&player.GB,
			&player.HRFB,
			&player.VFA,
			&player.ERA,
			&player.XERA,
			&player.FIP,
			&player.XFIP,
			&player.WAR,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

// GetPlayerByID will return a player
//
// it will query the position_players table based on a given player id
func (pool *DBPool) GetPitcherByID(id int) (*models.Pitcher, error) {
	query := `SELECT * FROM pitchers WHERE player_id = $1`
	player := &models.Pitcher{}

	err := pool.Poolconn.QueryRow(context.Background(), query, id).Scan(
		&player.ID,
		&player.Name,
		&player.Team,
		&player.W,
		&player.L,
		&player.SV,
		&player.G,
		&player.GS,
		&player.IP,
		&player.K9,
		&player.BB9,
		&player.HR9,
		&player.BABIP,
		&player.LOB,
		&player.GB,
		&player.HRFB,
		&player.VFA,
		&player.ERA,
		&player.XERA,
		&player.FIP,
		&player.XFIP,
		&player.WAR,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return player, nil
}

// AddPlayer will add a player to the position_players table
//
// will return nil if successful, error if unsuccessful
func (pool *DBPool) AddPitcher(player *models.Pitcher) error {
	query := `INSERT INTO pitchers 
	(name, team, w, l, sv, g, gs, ip, k9, bb9, hr9, babip, lob, gb, hrfb, vfa, era, xera, fip, xfip, war)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
	ON CONFLICT (name, team) DO NOTHING`

	_, err := pool.Poolconn.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.W,
		&player.L,
		&player.SV,
		&player.G,
		&player.GS,
		&player.IP,
		&player.K9,
		&player.BB9,
		&player.HR9,
		&player.BABIP,
		&player.LOB,
		&player.GB,
		&player.HRFB,
		&player.VFA,
		&player.ERA,
		&player.XERA,
		&player.FIP,
		&player.XFIP,
		&player.WAR,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePlayer will update a player in the position_players table given a player id
//
// if unsuccessful, will return an error
func (pool *DBPool) UpdatePitcher(player *models.Pitcher) error {
	query := `UPDATE pitchers SET
	name = $1,
	team = $2,
	w = $3,
	l = $4,
	sv = $5,
	g = $6,
	gs = $7,
	ip = $8,
	k9 = $9,
	bb9 = $10,
	hr9 = $11,
	babip = $12,
	lob = $13,
	gb = $14,
	hrfb = $15,
	vfa = $16,
	era = $17,
	xera = $18,
	fip = $19,
	xfip = $20,
	war = $21,
	WHERE player_id = $22`

	res, err := pool.Poolconn.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.W,
		&player.L,
		&player.SV,
		&player.G,
		&player.GS,
		&player.IP,
		&player.K9,
		&player.BB9,
		&player.HR9,
		&player.BABIP,
		&player.LOB,
		&player.GB,
		&player.HRFB,
		&player.VFA,
		&player.ERA,
		&player.XERA,
		&player.FIP,
		&player.XFIP,
		&player.WAR,
		&player.ID,
	)
	if err != nil {
		return err
	}

	log.Println("rows affected:", res.RowsAffected())

	return nil
}

// DeletePlayer deletes a player by player id in the position_players table
func (pool *DBPool) DeletePitcher(id int) error {
	query := `DELETE FROM pitchers WHERE player_id = $1`

	_, err := pool.Poolconn.Exec(context.Background(), query, id)
	return err
}
