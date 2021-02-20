package cmd

import (
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

func ProcessArgs(args []string) error {

	switch args[0] {
	case "db":
		parseDbFlags(args[1:])
	case "help":
		PrintHelp()
		os.Exit(0)
	}

	return nil
}
