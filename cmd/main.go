package main

import (
	"FSC/internal/cli"
	"FSC/internal/commands"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	commandsList := commands.CommandsAvailable()

	cli.InitCli()

	for index := range commandsList {
		commandsList[index].InitVariables()
	}
	flag.Parse()

	if len(os.Args) == 1 {

		fmt.Println("Please, provide a command to execute")

		cli.PrintMessage("Available commands: ")
		for index := range commandsList {
			command := commandsList[index]
			cli.PrintMessage(command.CommandName())
		}

		os.Exit(1)

	}

	var hadCommand bool = false
	for index := range commandsList {
		command := commandsList[index]
		mainArg := os.Args[1]
		if strings.Contains(mainArg, command.CommandName()) {
			hadCommand = true
			fmt.Println("Starting Command: " + command.CommandName())
			if cli.IsHelp {
				command.OnHelp()
			} else {
				command.Execute()
			}
		}

	}

	if !hadCommand {
		cli.PrintMessage("comando não encontrado")
	}
}
