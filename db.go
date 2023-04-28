package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// creates the DB type for handlers.go to utilize in the various routes
type DB interface {
	AddPlayer(*Player) error
	DeletePlayer(int) error
	UpdatePlayer(*Player) error
	GetPlayers() ([]*Player, error)
	GetPlayerByID(int) (*Player, error)
}

// Holds the pgxpool.Pool type for the initialization of the Postgres database via the pgx driver
type DBPool struct {
	db *pgxpool.Pool
}

// NewDBPool returns a DBPool instance
//
// sets up a new Postgres pgx pool
// if pgx pool is unable to be set up, will fail
func NewDBPool() (*DBPool, error) {
	godotenv.Load()
	db_url := os.Getenv("POSTGRES_URL")

	dbpool, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}

	log.Println("Connected to db on port:", os.Getenv("PORT"))

	return &DBPool{
		db: dbpool,
	}, nil
}

// InitializePlayerTable will create the position_players postgres table upon initialization
func (pool *DBPool) InitializePlayerTable() error {
	return pool.CreatePlayerTable()
}

// CreatePlayerTable executes the query to create the position_player table with pgx
func (pool *DBPool) CreatePlayerTable() error {
	query := `create table if not exists position_players (
		player_id serial primary key NOT NULL,
		name varchar(50) NOT NULL,
		team varchar(3),
		games int CHECK (games >= 0),
		pa int CHECK (pa >= 0),
		hr int CHECK (hr >= 0),
		runs int CHECK (runs >= 0),
		rbi int CHECK (rbi >= 0),
		sb int CHECK (sb >= 0),
		bb_rate float(3) CHECK (bb_rate >= 0),
		k_rate float(3) CHECK (k_rate >= 0),
		iso float(3) CHECK (iso >= 0),
		average float(3) CHECK (average >= 0),
		obp float(3) CHECK (obp >= 0),
		slg float(3) CHECK (slg >= 0),
		woba float(3) CHECK (woba >= 0),
		unique (name, team)
	)`

	_, err := pool.db.Exec(context.Background(), query)

	return err
}

// ClearPlayerTable clears the position_players table for testing purposes
func (pool *DBPool) ClearPlayerTable() error {
	clear_records_query := `DROP * FROM position_players`
	reset_id_query := `ALTER SEQUENCE id RESTART WITH 1`

	_, err_clear := pool.db.Exec(context.Background(), clear_records_query)
	if err_clear != nil {
		return err_clear
	}
	_, err_reset := pool.db.Exec(context.Background(), reset_id_query)
	if err_reset != nil {
		return err_reset
	}

	return nil
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

func ReadFromCSV() []*Player {
	file, err := os.Open("/usr/src/baseball_api/stats.csv")
	if err != nil {
		log.Fatalf("Unable to open CSV file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read CSV file: %v\n", err)
	}

	var players []*Player

	for i, record := range records {
		if i == 0 {
			continue
		}

		row := &Player{
			Name:   record[0],
			Team:   record[1],
			Games:  ConvertToInt(record[2]),
			PA:     ConvertToInt(record[3]),
			HR:     ConvertToInt(record[4]),
			R:      ConvertToInt(record[5]),
			RBI:    ConvertToInt(record[6]),
			SB:     ConvertToInt(record[7]),
			BbRate: ConvertToFloat(record[8]),
			KRate:  ConvertToFloat(record[9]),
			ISO:    ConvertToFloat(record[10]),
			AVG:    ConvertToFloat(record[11]),
			OBP:    ConvertToFloat(record[12]),
			SLG:    ConvertToFloat(record[13]),
			WOBA:   ConvertToFloat(record[14]),
		}
		players = append(players, row)
	}

	return players
}

// GetPlayers will return a list of players
//
// retrieves all existing players in the position_players table
func (pool *DBPool) GetPlayers() ([]*Player, error) {
	query := `SELECT * FROM position_players`

	rows, err := pool.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	players := []*Player{}
	for rows.Next() {
		player := &Player{}
		err := rows.Scan(
			&player.ID,
			&player.Name,
			&player.Team,
			&player.Games,
			&player.PA,
			&player.HR,
			&player.R,
			&player.RBI,
			&player.SB,
			&player.BbRate,
			&player.KRate,
			&player.ISO,
			&player.AVG,
			&player.OBP,
			&player.SLG,
			&player.WOBA,
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
func (pool *DBPool) GetPlayerByID(id int) (*Player, error) {
	query := `SELECT * FROM position_players WHERE player_id = $1`
	player := &Player{}

	err := pool.db.QueryRow(context.Background(), query, id).Scan(
		&player.ID,
		&player.Name,
		&player.Team,
		&player.Games,
		&player.PA,
		&player.HR,
		&player.R,
		&player.RBI,
		&player.SB,
		&player.BbRate,
		&player.KRate,
		&player.ISO,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
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
func (pool *DBPool) AddPlayer(player *Player) error {
	query := `INSERT INTO position_players 
	(name, team, games, pa, hr, runs, rbi, sb, bb_rate, k_rate, iso, average, obp, slg, woba)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	ON CONFLICT (name, team) DO NOTHING`

	_, err := pool.db.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.Games,
		&player.PA,
		&player.HR,
		&player.R,
		&player.RBI,
		&player.SB,
		&player.BbRate,
		&player.KRate,
		&player.ISO,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePlayer will update a player in the position_players table given a player id
//
// if unsuccessful, will return an error
func (pool *DBPool) UpdatePlayer(player *Player) error {
	query := `UPDATE position_players SET
	name = $1,
	team = $2,
	games = $3,
	pa = $4,
	hr = $5,
	runs = $6,
	rbi = $7,
	sb = $8,
	bb_rate = $9,
	k_rate = $10,
	iso = $11,
	average = $12,
	obp = $13,
	slg = $14,
	woba = $15,
	WHERE player_id = $16`

	res, err := pool.db.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.Games,
		&player.PA,
		&player.HR,
		&player.R,
		&player.RBI,
		&player.SB,
		&player.BbRate,
		&player.KRate,
		&player.ISO,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
		&player.ID,
	)
	if err != nil {
		return err
	}

	log.Println("rows affected:", res.RowsAffected())

	return nil
}

// DeletePlayer deletes a player by player id in the position_players table
func (pool *DBPool) DeletePlayer(id int) error {
	query := `DELETE FROM position_players WHERE player_id = $1`

	_, err := pool.db.Exec(context.Background(), query, id)
	return err
}
