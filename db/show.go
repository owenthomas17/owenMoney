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

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var expense_name string
		var expense_type string
		var cost float64

		err = rows.Scan(&id, &expense_name, &expense_type, &cost)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, expense_name, expense_type, cost)
	}

}
