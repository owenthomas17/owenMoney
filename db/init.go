package db

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"owenMoney/constants"
)

type column struct {
	name   				string
	dbType 				string
	pk     				bool
	autoincrement bool
}

type table struct {
	columns []column
	tblName string
}

func (c *column) columnStringBuilder() string {

	if c.pk == true {
		if c.autoincrement == true {
			return fmt.Sprintf("%s %s not null primary key autoincrement", c.name, c.dbType)
		}
		return fmt.Sprintf("%s %s not null primary key", c.name, c.dbType)
	}

	return fmt.Sprintf("%s %s", c.name, c.dbType)
}

func InitDb() {
	if cleanDb() {
		createDbTables()
	}
}

func createDbTables() {
	log.Printf("Connecting to database at: %s", constants.DbFullFilePath)
	db, err := sql.Open("sqlite3", constants.DbFullFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	id := column{"id", "integer", true, true}
	expenseName := column{"expense_name", "text", false, false}
	expenseType := column{"expense_type", "text", false, false}
	cost := column{"cost", "real", false, false}

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
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

}

func checkDbFileExists() bool {
	log.Printf("Checking if the database file already exists at: %s", constants.DbFullFilePath)
	info, err := os.Stat(constants.DbFullFilePath)

	if err != nil {
		log.Printf("Database file doesn't already exist")
		return false
	}

	return !info.IsDir()
}

func cleanDb() bool {
	if checkDbFileExists() {
		log.Printf("Database file already exists, cleaning by removing file: %s", constants.DbFullFilePath)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Removing database file, this will delete all contents. Are you sure? y/N: ")
		confirmation, _ := reader.ReadString('\n')
		if confirmation == "y\n" || confirmation == "Y\n" {
			log.Printf("Deleting database file: %s", constants.DbFullFilePath)
			os.Remove(constants.DbFullFilePath)
			return true
		} else {
			log.Printf("Leaving existing DB in place")
			return false
		}
	}
	return true
}
