package models

import "log"

type Player struct {
	PlayerId string `json:"playerId"`
	Points float64 `json:"balance"`
}

func (db *DB) SelectPlayer(playerId string) Player{
	rows, err := db.Query("select * from Players where playerId=?", playerId)
	if err != nil {
		log.Fatal(err)
	}
	var selectedPlayer Player

	for rows.Next() {
		rows.Scan(&selectedPlayer.PlayerId, &selectedPlayer.Points)
	}
	return selectedPlayer
}
func (db *DB) TakeFromPlayer(playerId string, points float64){
	_, err := db.Exec("update Players set points = points - ? where playerId=?", points, playerId)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *DB) CreateNewPlayer(playerId string, points float64){
	_, err := db.Exec("insert into Players values (?, ?)", playerId, points)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *DB) FundInPlayer(playerId string, points float64){
	_, err := db.Exec("update Players set points = points + ? where playerId=?", points, playerId)
	if err != nil {
		log.Fatal(err)
	}
}
