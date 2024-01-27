package commands

type Command interface {
	IsMatchCommand() bool
	Execute()
	OnHelp()
}

