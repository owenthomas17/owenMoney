package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"owenMoney/constants"
)



func ShowAll() {
	log.Print(constants.DbFullFilePath)
	db, err := sql.Open("sqlite3", constants.DbFullFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Printf("Showing all rows in database")
	sqlStmt := fmt.Sprintf(`
	    select * from tbl_default
	`)

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

}
