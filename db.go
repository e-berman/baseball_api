package main

import (
	"fmt"
	"context"
	"os"
	"log"

	"github.com/joho/godotenv"
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
	godotenv.Load()

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	} else {
		fmt.Printf("Connected to db on port: %s\n", os.Getenv("PORT"))
	}
	
	return &DBPool{
		db: dbpool,
	}, nil
}

func (pool *DBPool) InitializePlayerTable() error {
	return pool.createPlayerTable()
}

func (pool *DBPool) createPlayerTable() error {
	query := `create table if not exists position_players (
		player_id serial primary key NOT NULL,
		player_name varchar(50) NOT NULL,
		team varchar(3),
		position varchar(10) NOT NULL,
		games integer CHECK (games >= 0),
		pa integer CHECK (pa >= 0),
		hr integer CHECK (hr >= 0),
		runs integer CHECK (runs >= 0),
		rbi integer CHECK (rbi >= 0),
		sb integer CHECK (sb >= 0),
		bb_rate float(3) CHECK (bb_rate >= 0),
		k_rate float(3) CHECK (k_rate >= 0),
		iso float(3) CHECK (iso >= 0),
		babip float(3) CHECK (babip >= 0),
		average float(3) CHECK (average >= 0),
		obp float(3) CHECK (obp >= 0),
		slg float(3) CHECK (slg >= 0),
		woba float(3) CHECK (woba >= 0),
		wrc_plus integer CHECK (wrc_plus >= 0),
		war float(3),
		unique (player_name, team, war)
	)`

	_, err := pool.db.Exec(context.Background(), query)
	return err
}

func (pool *DBPool) AddPlayer(player *Player) error {
	query := `insert into position_players 
	(player_name, team, position, games, pa, hr, runs, rbi, sb, wrc_plus, bb_rate, k_rate, iso, babip, average, obp, slg, woba, war)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
	on conflict (player_name, team, war) do nothing`

	_, err := pool.db.Exec(context.Background(), query,
		player.PlayerName,
		player.Team,
		player.Position,
		player.Games,
		player.PA,
		player.HR,
		player.R,
		player.RBI,
		player.SB,
		player.WRCPlus,
		player.BbRate,
		player.KRate,
		player.ISO,
		player.BABIP,
		player.AVG,
		player.OBP,
		player.SLG,
		player.WOBA,
		player.LastSeasonWAR,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pool *DBPool) DeletePlayer(id int) error {
	query := `delete from position_players where player_id = $1`

	_, err := pool.db.Exec(context.Background(), query, id)
	return err
}

func (pool *DBPool) UpdatePlayer(*Player) error {
	return nil
}

func (pool *DBPool) GetPlayers() ([]*Player, error) {
	query := `select * from position_players`

	rows, err := pool.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	players := []*Player{}
	for rows.Next() {
		player := &Player{}
		err := rows.Scan(
			&player.ID,
			&player.PlayerName,
			&player.Team,
			&player.Position,
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
	query := `select * from position_players where player_id = $1`
	player := &Player{}

	err := pool.db.QueryRow(context.Background(), query, id).Scan(
		&player.ID,
		&player.PlayerName,
		&player.Team,
		&player.Position,
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
