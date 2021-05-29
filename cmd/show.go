package cmd

import (
	"os"
	"flag"
	"fmt"
	"log"
	"owenMoney/db"
	"owenMoney/constants"
)


func printShowHelp() {
	fmt.Println(`
Please choose one of the following commands:

owenmoney show all   - Shows all expenses
owenmoney show help  - Prints this help message
    `)
}

func showInitFlagSet() *showInitFlags {

	showFlags := &showInitFlags{
		fs: flag.NewFlagSet("show", flag.ExitOnError),
	}

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	showFlags.fs.StringVar(&showFlags.filterColumn, "filter-column", "default", "filter column")
	showFlags.fs.StringVar(&showFlags.filterString, "filter-string", "default", "filter filter string used on column")
	showFlags.fs.StringVar(&showFlags.dbName, "db-name", "default", "database file to read from")
	showFlags.fs.StringVar(&showFlags.dbFilePath, "db-file-path", workingDir, "database file to read from")

	return showFlags
}

type showInitFlags struct {
	fs         *flag.FlagSet
	filterColumn     string
	filterString     string
	dbName			     string
	dbFilePath	     string
}

func parseShowFlags(args []string) {
	if len(args) < 1 {
		printShowHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "all":
		showFlagSet := showInitFlagSet()
		if err := showFlagSet.fs.Parse(args[1:]); err == nil {
			constants.SetDbFileName(showFlagSet.dbName)
			constants.SetDbFilePath(showFlagSet.dbFilePath)
			constants.SetDbFullFilePath()
			db.ShowAll()
		}
	default:
		printShowHelp()
	}
}
