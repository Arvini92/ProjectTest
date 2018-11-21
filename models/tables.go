package models

import "log"

func (db *DB) CreateTables(){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Backers (
			playerId TEXT NOT NULL,
			backerId TEXT NOT NULL,
			tournamentId TEXT NOT NULL,
			amount REAL NOT NULL
		);
		CREATE TABLE IF NOT EXISTS Players (
			playerId TEXT NOT NULL UNIQUE,
			points REAL NOT NULL
		);
		CREATE TABLE IF NOT EXISTS Tournaments(
			tournamentId TEXT NOT NULL UNIQUE,
			deposit	REAL NOT NULL
		);
	`
	_, err := db.Exec(sql_table)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) DropTables(){
	_, err := db.Exec("DROP TABLE Tournaments")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DROP TABLE Players")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DROP TABLE Backers")
	if err != nil {
		log.Fatal(err)
	}
}