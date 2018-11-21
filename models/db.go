package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

type Datastore interface {
	SelectPlayer(playerId string) Player
	TakeFromPlayer(playerId string, points float64)
	CreateNewPlayer(playerId string, points float64)
	FundInPlayer(playerId string, points float64)
	AddTournament (tournamentId string, deposit float64)
	SelectTournament(tournamentId string) Tournament
	CreateNewBacker(playerId string, backerId string, tournamentId string, amount float64)
	PlayerInGame (playerId string, tournamentId string) bool
	SelectBackers(winnerId string) []string
	ZeroBackers(tournamentId string)
	CreateTables()
	DropTables()
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}


