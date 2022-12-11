package db

import (
	"database/sql"
	"fmt"
	"log"
)

var Database *sql.DB

func GetMySQLConnection(DBUser, DBPass, DBHost, DBBase string) {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBBase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Fatal(err)
	}
	Database = db
}
