package config

import (
	"github.com/jmcvetta/neoism"
)

var db *neoism.Database

func CreateDBConnection() {
	conn, err := neoism.Connect("http://neo4j:test@localhost:7474/db/data")
	if err != nil {
		panic(err)
	} else {
		db = conn
	}
}

func GetDBConnection() *neoism.Database {
	return db
}
