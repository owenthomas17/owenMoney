package main

import (
	"flag"
	"fmt"
	"os"
)

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

func printHelp() {
	fmt.Println(`
Please choose one of the following commands:
db
help
    `)

}

func processArgs(args []string) error {

	switch args[0] {
	case "db":
		dbFlagSet := DbFlagSet()
		if err := dbFlagSet.fs.Parse(args[1:]); err == nil {
			fmt.Println(dbFlagSet.dbName)
			fmt.Println(dbFlagSet.tblName)
		}
	case "help":
		printHelp()
		os.Exit(0)
	}

	return nil
}

func main() {
	args := os.Args

	if len(args) <= 1 {
		printHelp()
		os.Exit(1)
	}

	processArgs(args[1:])
}
