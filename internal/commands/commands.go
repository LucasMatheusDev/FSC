package commands

func CommandsAvailable() []CommandInterface {
	return []CommandInterface{
		CleanDartArch{},
		CreateTest{},
		ConfigCommand{},
	}
}
