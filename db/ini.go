package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type column struct {
	name   string
	dbType string
	pk     bool
}

type table struct {
	column  *column
	tblName string
}

func CreateDB(f string) {
	filePath := fmt.Sprintf("./%s.db", f)

	os.Remove(filePath)

	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := fmt.Sprintf("create table tbl_default (id integer not null primary key, name text);")

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

}
