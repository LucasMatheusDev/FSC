package commands

import (
	"FSC/internal/cli"
	"FSC/internal/commands/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CleanDartArch struct {
}

func (c CleanDartArch) IsMatchCommand() bool {
	moduleName = ""
	flag.StringVar(&moduleName, "create-module", "", "Module name to create")
	flag.Parse()
	return moduleName != ""

}

func (c CleanDartArch) Execute() {
	createCleanDart()
}

func (c CleanDartArch) OnHelp() {
	cli.PrintMessage("Comando para criar a arquitetura Clean dart")
}

var moduleName string

func createCleanDart() {

	if moduleName == "" {
		fmt.Println("Por favor, forneça o nome do módulo usando a flag -create-module")
		os.Exit(1)
	}

	var currentDir string

	currentDir = utils.GetProjectPath()
	if currentDir == "" {
		var err error = nil
		currentDir, err = os.Getwd()
		if err != nil {
			fmt.Printf("Erro ao obter o diretório de trabalho: %s\n", err)
			os.Exit(1)
		}
	}

	directories := []string{
		currentDir + "/" + moduleName + "/" + moduleName + "_module.dart",
		currentDir + "/" + moduleName + "/" + moduleName + "_routes.dart",

		/// Domain
		currentDir + "/" + moduleName + "/" + "domain" + "/" + "entities" + "/" + moduleName + "_entity.dart",
		currentDir + "/" + moduleName + "/" + "domain" + "/" + "usecases" + "/" + moduleName + "_usecase.dart",
		currentDir + "/" + moduleName + "/" + "domain" + "/" + "repositories" + "/" + moduleName + "_repository.dart",

		/// Infra
		currentDir + "/" + moduleName + "/" + "infra" + "/" + "repositories" + "/" + moduleName + "_repository_impl.dart",
		currentDir + "/" + moduleName + "/" + "infra" + "/" + "models" + "/" + moduleName + "_model.dart",
		currentDir + "/" + moduleName + "/" + "infra" + "/" + "data" + "/" + "data_sources" + "/" + moduleName + "_data_source.dart",

		/// Data
		currentDir + "/" + moduleName + "/" + "external" + "/" + "data" + "/" + "data_sources" + "/" + moduleName + "_data_source_impl.dart",

		/// Presenter
		currentDir + "/" + moduleName + "/" + "presenter" + "/" + "controllers" + "/" + moduleName + "_controller.dart",
		currentDir + "/" + moduleName + "/" + "presenter" + "/" + "view" + "/" + "pages" + "/" + moduleName + "_page.dart",
		currentDir + "/" + moduleName + "/" + "presenter" + "/" + "view" + "/" + "delegates" + "/" + moduleName + "_delegate.dart",
		currentDir + "/" + moduleName + "/" + "presenter" + "/" + "view" + "/" + "params/",
		currentDir + "/" + moduleName + "/" + "presenter" + "/" + "view" + "/" + "widgets/",
	}

	for _, dir := range directories {
		dir = strings.ReplaceAll(dir, "\\", "/")

		if err := os.MkdirAll(filepath.Dir(dir), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if strings.Contains(dir, ".dart") {
			file, err := os.Create(dir)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}

		cli.PrintVerboseMessage("Criado com sucesso: " + dir)
	}
	cli.PrintMessage("Módulo" + moduleName + " criado com sucesso!")
}
