package models

import "log"

type Tournament struct {
	TournamentId string
	Deposit float64
}

func (db *DB) AddTournament (tournamentId string, deposit float64){
	_, err := db.Exec("insert into Tournaments values (?, ?)", tournamentId, deposit)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *DB) SelectTournament(tournamentId string) Tournament{
	rows, err := db.Query("select * from Tournaments where tournamentId=?", tournamentId)
	if err != nil {
		log.Fatal(err)
	}
	var tournament Tournament
	for rows.Next() {
		rows.Scan(&tournament.TournamentId, &tournament.Deposit)
	}
	return tournament
}
