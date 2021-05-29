package cmd

import (
	"os"
	"flag"
	"fmt"
	"log"
	"owenMoney/db"
	"owenMoney/constants"
)


func printAddHelp() {
	fmt.Println(`
Please choose one of the following commands:

owenmoney add expense - Add an expense to the database
owenmoney add help    - Prints this help message
    `)
}

func addInitFlagSet() *addInitFlags {

	addFlags := &addInitFlags{
		fs: flag.NewFlagSet("show", flag.ExitOnError),
	}

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	addFlags.fs.StringVar(&addFlags.dbName, "db-name", "default", "database file to read from")
	addFlags.fs.StringVar(&addFlags.dbFilePath, "db-file-path", workingDir, "database file to read from")
	addFlags.fs.StringVar(&addFlags.expenseName, "expense-name", "invalid", "Name of the expense, must be set")
	addFlags.fs.StringVar(&addFlags.expenseType, "expense-type", "invalid", "Type of expense, must be set")
	addFlags.fs.Float64Var(&addFlags.cost, "cost", 0.00, "Cost in pounds")

	return addFlags
}

type addInitFlags struct {
	fs         *flag.FlagSet
	dbName			     string
	dbFilePath	     string
	expenseName      string
	expenseType      string
	cost             float64
}

func parseAddFlags(args []string) {
	if len(args) < 1 {
		printAddHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "expense":
		showFlagSet := addInitFlagSet()
		if err := showFlagSet.fs.Parse(args[1:]); err == nil {
			constants.SetDbFileName(showFlagSet.dbName)
			constants.SetDbFilePath(showFlagSet.dbFilePath)
			constants.SetDbFullFilePath()
			db.AddExpense(showFlagSet.expenseName, showFlagSet.expenseType, showFlagSet.cost)
		}
	default:
		printAddHelp()
	}
}
