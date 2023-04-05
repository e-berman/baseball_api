package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	AddPlayer(*Player) error
	DeletePlayer(int) error
	UpdatePlayer(*Player) error
	GetPlayers() ([]*Player, error)
	GetPlayerByID(int) (*Player, error)
}

type DBPool struct {
	db *pgxpool.Pool
}

func NewDBPool() (*DBPool, error) {
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

func (pool *DBPool) InitializePlayerTable() error {
	return pool.CreatePlayerTable()
}

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
		babip float(3) CHECK (babip >= 0),
		average float(3) CHECK (average >= 0),
		obp float(3) CHECK (obp >= 0),
		slg float(3) CHECK (slg >= 0),
		woba float(3) CHECK (woba >= 0),
		wrc_plus int CHECK (wrc_plus >= 0),
		war float(3),
		unique (name, team, war)
	)`

	_, err := pool.db.Exec(context.Background(), query)

	return err
}

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
			&player.BABIP,
			&player.AVG,
			&player.OBP,
			&player.SLG,
			&player.WOBA,
			&player.WRCPlus,
			&player.LastSeasonWAR,
		) 
			
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

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
		&player.BABIP,
		&player.AVG,
		&player.OBP,
		&player.SLG,
		&player.WOBA,
		&player.WRCPlus,
		&player.LastSeasonWAR,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return player, nil
}

func (pool *DBPool) AddPlayer(player *Player) error {
	query := `INSERT INTO position_players 
	(name, team, games, pa, hr, runs, rbi, sb, wrc_plus, bb_rate, k_rate, iso, babip, average, obp, slg, woba, war)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	ON CONFLICT (name, team, war) DO NOTHING`

	_, err := pool.db.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.Games,
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
		&player.LastSeasonWAR,
	)
	if err != nil {
		return err
	}

	return nil
}

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
	wrc_plus = $9,
	bb_rate = $10,
	k_rate = $11,
	iso = $12,
	babip = $13,
	average = $14,
	obp = $15,
	slg = $16,
	woba = $17,
	war = $18
	WHERE player_id = $19`

	res, err := pool.db.Exec(context.Background(), query,
		&player.Name,
		&player.Team,
		&player.Games,
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
		&player.LastSeasonWAR,
		&player.ID,
	)
	if err != nil {
		return err
	}

	log.Println("rows affected:", res.RowsAffected())

	return nil
}

func (pool *DBPool) DeletePlayer(id int) error {
	query := `DELETE FROM position_players WHERE player_id = $1`

	_, err := pool.db.Exec(context.Background(), query, id)
	return err
}


