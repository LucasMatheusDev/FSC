package main

import (
	"FSC/internal/commands"
	"FSC/internal/cli"
	"fmt"
	"os"
	"strings"
)

func main() {
	cli.InitCli()
	var firstArg string

	if len(os.Args) > 1 {
		firstArg = os.Args[1]
	} else {
		fmt.Println("Por favor, forneça o nome do módulo usando a flag -module")
		os.Exit(1)
	}

	if strings.Contains(firstArg, "create-module") {
		commands.CreateCleanDart()
	} else {
		fmt.Println("Por favor, forneça o nome do módulo usando a flag -module")
		os.Exit(1)
	}
}
