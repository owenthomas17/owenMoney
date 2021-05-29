package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"owenMoney/constants"
	"owenMoney/db"
)

func printDbHelp() {
	fmt.Println(`
Please choose one of the following commands:

owenmoney db init  - Command set to initialise the database
owenmoney db help  - Prints this help message
    `)

}

func DbInitFlagSet() *dbInitFlags {

	dbFlags := &dbInitFlags{
		fs: flag.NewFlagSet("init", flag.ExitOnError),
	}

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbFlags.fs.StringVar(&dbFlags.dbName, "db-name", "default", "default DB filename")
	dbFlags.fs.StringVar(&dbFlags.dbFilePath, "db-file-path", workingDir, "default DB filepath")

	return dbFlags
}

type dbInitFlags struct {
	fs         *flag.FlagSet
	dbName     string
	dbFilePath string
}

func parseDbFlags(args []string) error {

	if len(args) < 1 {
		printDbHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "init":
		dbFlagSet := DbInitFlagSet()
		if err := dbFlagSet.fs.Parse(args[1:]); err == nil {
			constants.SetDbFileName(dbFlagSet.dbName)
			constants.SetDbFilePath(dbFlagSet.dbFilePath)
			constants.SetDbFullFilePath()
			db.InitDb()
		}
	default:
		printDbHelp()
	}
	return nil
}
