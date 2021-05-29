package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"owenMoney/constants"
)

type column struct {
	name   string
	dbType string
	pk     bool
}

type table struct {
	columns []column
	tblName string
}

func (c *column) columnStringBuilder() string {

	if c.pk == true {
		return fmt.Sprintf("%s %s not null primary key", c.name, c.dbType)
	}

	return fmt.Sprintf("%s %s", c.name, c.dbType)
}

func InitDb() {
	fmt.Println(constants.DbFilePath)
	createDb()
	createDbTables()

}

func createDbTables() {
	log.Printf("Connecting to database at: %s", constants.DbFullFilePath)
	db, err := sql.Open("sqlite3", constants.DbFullFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	id := column{"id", "integer", true}
	expenseName := column{"expense_name", "text", false}
	expenseType := column{"expense_type", "text", false}
	cost := column{"cost", "integer", false}

	log.Printf("Creating the default tables for %s", constants.DbFileName)
	sqlStmt := fmt.Sprintf(`
	    CREATE TABLE tbl_default (
		%s,
		%s,
		%s,
		%s);
	    `, id.columnStringBuilder(),
				 expenseName.columnStringBuilder(),
				 expenseType.columnStringBuilder(),
				 cost.columnStringBuilder())

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

}

func checkDbFileExists() bool {
	log.Printf("Checking if the database file already exists at: %s", constants.DbFullFilePath)
	info, err := os.Stat(constants.DbFullFilePath)

	if err != nil {
		return false
	}

	return !info.IsDir()
}

func createDb() {
	checkDbFileExists()
	log.Printf("Removing current database file at: %s", constants.DbFullFilePath)
	os.Remove(constants.DbFullFilePath)

}
