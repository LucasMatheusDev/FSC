package cli

import (
	"flag"
)

var IsVerboseEnable = false

var IsHelp bool = false

func InitCli() {
	IsVerboseEnable = true
	IsHelp = false
	checkVerboseEnable()
	checkIsHelp()

}

func checkVerboseEnable() {
	flag.BoolVar(&IsVerboseEnable, "v", false, "Ativar modo verbose")
}

func checkIsHelp() {
	flag.BoolVar(&IsHelp, "h", false, "Ask Help")
}

func PrintVerboseMessage(message string) {
	if IsVerboseEnable {
		println(message)
	}
}

func PrintMessage(message string) {
	println(message)
}
