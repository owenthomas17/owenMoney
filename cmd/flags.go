package cmd

import (
	"fmt"
	"os"
)

func PrintHelp() {
	fmt.Println(`
Please choose one of the following commands:

owenmoney db   - Interacts configuration of the backend database
owenmoney help - Prints this help message
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
