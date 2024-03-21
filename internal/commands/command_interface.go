package commands

type CommandInterface interface {
	InitVariables() 
	Execute()
	OnHelp()
	CommandName() string

}
