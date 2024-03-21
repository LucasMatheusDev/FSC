package commands

import (
	"FSC/internal/cli"
	"FSC/internal/commands/utils"
	"flag"
	"os"
)

type ConfigCommand struct {
}

func (c ConfigCommand) CommandName() string {
	return "config"
}

var projectPath string = ""

func (c ConfigCommand) InitVariables() {
	flag.Bool(c.CommandName(), false, "Configuration of FSC")
	flag.StringVar(&projectPath, "project-path", "", "Default Project Path")
}

func (c ConfigCommand) Execute() {
	config()
}

func (c ConfigCommand) OnHelp() {
	cli.PrintMessage("Command to config FSC")
}

func config() {
	if len(os.Args) < 2 {
		cli.PrintMessage("Please, provide a command with flag to execute")
		return
	}

	if projectPath != "" {
		defineProjectPath()
	}
}

func defineProjectPath() {
	_, err := utils.SetProjectPath(projectPath)

	if err != nil {
		cli.PrintMessage("failed to add to PATH: " + err.Error())
	} else {
		cli.PrintMessage("Added Default Project Path: " + projectPath)
	}
}
