package cmd

import (
	"fmt"
	"os"
)

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
