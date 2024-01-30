package main

import (
	"FSC/internal/cli"
	"FSC/internal/commands"
	"fmt"
	"os"
)

func main() {

	commandsList := commands.CommandsAvailable()

	cli.InitCli()

	if len(os.Args) == 1 {

		fmt.Println("Por favor, forneça o nome do módulo usando a flag -module")
		os.Exit(1)

	}

	for command := range commandsList {
		if commandsList[command].IsMatchCommand() {
			if cli.IsHelp {
				commandsList[command].OnHelp()
			} else {
				commandsList[command].Execute()
			}
		} else {
			cli.PrintMessage("comando não encontrado")
		}

	}
}
