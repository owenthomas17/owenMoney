package cmd

import (
	"flag"
	"fmt"
	"os"
)

func PrintHelp() {
	fmt.Println(`
Please choose one of the following commands:
db
help
    `)

}

func DbFlagSet() *DbFlags {

	dbFlags := &DbFlags{
		fs: flag.NewFlagSet("db", flag.ExitOnError),
	}

	dbFlags.fs.StringVar(&dbFlags.dbName, "name", "default", "default DB filename")
	dbFlags.fs.StringVar(&dbFlags.tblName, "table", "tbl_default", "default table name")

	return dbFlags
}

type DbFlags struct {
	fs      *flag.FlagSet
	dbName  string
	tblName string
}

func ProcessArgs(args []string) error {

	switch args[0] {
	case "db":
		dbFlagSet := DbFlagSet()
		if err := dbFlagSet.fs.Parse(args[1:]); err == nil {
			fmt.Println(dbFlagSet.dbName)
			fmt.Println(dbFlagSet.tblName)
		}
	case "help":
		PrintHelp()
		os.Exit(0)
	}

	return nil
}
