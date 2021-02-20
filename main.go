package main

import (
	"os"
	"owenMoney/cmd"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		cmd.PrintHelp()
		os.Exit(0)
	}

	cmd.ProcessArgs(args[1:])
}
