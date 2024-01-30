package commands

type CommandInterface interface {
	IsMatchCommand() bool
	Execute()
	OnHelp()
}
