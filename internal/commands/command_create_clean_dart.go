package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateCleanDart() {
	var moduleName string
	flag.StringVar(&moduleName, "module", "", "Nome do módulo a ser criado")
	flag.Parse()

	// Verifique se o nome do módulo foi fornecido
	if moduleName == "" {
		fmt.Println("Por favor, forneça o nome do módulo usando a flag -module")
		os.Exit(1)
	}

	// Obtém o diretório atual de trabalho
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erro ao obter o diretório de trabalho: %s\n", err)
		os.Exit(1)
	}

	// Lista de diretórios para criar
	directories := []string{
		/// Module
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
		// Salva pastas e arquivos
		dir = strings.ReplaceAll(dir, "\\", "/")

		// Cria os diretórios pais
		if err := os.MkdirAll(filepath.Dir(dir), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// Cria o arquivo
		if strings.Contains(dir, ".dart") {
			file, err := os.Create(dir)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}

		fmt.Println("Criado com sucesso:", dir)
	}
}
