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

func (c *column) columnString() string {

	if c.pk == true {
		return fmt.Sprintf("%s %s not null primary key", c.name, c.dbType)
	}

	return fmt.Sprintf("%s %s", c.name, c.dbType)
}

func CreateDB() {
	os.Remove(constants.DbFullFilePath)

	db, err := sql.Open("sqlite3", constants.DbFullFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	id := column{"id", "integer", true}
	name := column{"name", "text", false}

	sqlStmt := fmt.Sprintf(`
	    CREATE TABLE tbl_default (
		%s,
		%s);
	    `, id.columnString(), name.columnString())

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

}
