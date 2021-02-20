package cmd

import (
	"flag"
	"fmt"
)

func DbInitFlagSet() *dbInitFlags {

	dbFlags := &dbInitFlags{
		fs: flag.NewFlagSet("db", flag.ExitOnError),
	}

	dbFlags.fs.StringVar(&dbFlags.dbName, "name", "default", "default DB filename")
	dbFlags.fs.StringVar(&dbFlags.tblName, "table", "tbl_default", "default table name")

	return dbFlags
}

type dbInitFlags struct {
	fs      *flag.FlagSet
	dbName  string
	tblName string
}

func parseDbFlags(args []string) error {
	switch args[0] {
	case "init":
		dbFlagSet := DbInitFlagSet()
		if err := dbFlagSet.fs.Parse(args[1:]); err == nil {
			fmt.Println(dbFlagSet.dbName)
			fmt.Println(dbFlagSet.tblName)
		}
	}
	return nil
}
