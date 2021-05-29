package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"owenMoney/constants"
)

func AddExpense(name string, etype string, cost float64) {
	log.Print(constants.DbFullFilePath)
	db, err := sql.Open("sqlite3", constants.DbFullFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Printf("Inserting into database")
	sqlStmt := fmt.Sprintf(`
			insert into tbl_default (expense_name,expense_type,cost) 
			VALUES ("%s", "%s", %f);
	`, name, etype, cost)

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

}
