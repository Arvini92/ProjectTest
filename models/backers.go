package models

import "log"

func (db *DB)CreateNewBacker(playerId string, backerId string, tournamentId string, amount float64){
	_, err := db.Exec("insert into Backers values (?, ?, ?, ?)", playerId, backerId, tournamentId, amount)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB)PlayerInGame (playerId string, tournamentId string) bool{
	rows, err := db.Query("select playerId from Backers where playerId=backerId and playerId=? and tournamentId=?", playerId, tournamentId)
	if err != nil {
		log.Fatal(err)
	}
	var playInGameCheck string
	for rows.Next() {
		rows.Scan(&playInGameCheck)
	}
	if playInGameCheck == "" {
		return false
	} else {
		return true
	}
}

func (db *DB)SelectBackers(winnerId string) []string {
	rows, err := db.Query("select backerId from Backers where playerId=?", winnerId)
	if err != nil {
		log.Fatal(err)
	}
	var backerId string
	var backerIds[] string

	for rows.Next() {
		rows.Scan(&backerId)
		backerIds = append(backerIds, backerId)
	}
	return backerIds
}

func (db *DB) ZeroBackers(tournamentId string){
	_, err := db.Exec("update Backers set amount = 0 where tournamentId=?", tournamentId)
	if err != nil {
		log.Fatal(err)
	}
}
