package cmd

import (
	"flag"
	"os"
	"owenMoney/db"
)

func DbInitFlagSet() *dbInitFlags {

	dbFlags := &dbInitFlags{
		fs: flag.NewFlagSet("init", flag.ExitOnError),
	}

	dbFlags.fs.StringVar(&dbFlags.dbName, "name", "default", "default DB filename")

	return dbFlags
}

type dbInitFlags struct {
	fs     *flag.FlagSet
	dbName string
}

func parseDbFlags(args []string) error {

	if len(args) < 1 {
		PrintHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "init":
		dbFlagSet := DbInitFlagSet()
		if err := dbFlagSet.fs.Parse(args[1:]); err == nil {
			db.CreateDB(dbFlagSet.dbName)
		}
	default:
		PrintHelp()
	}
	return nil
}
