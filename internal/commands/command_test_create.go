package commands

import (
	"FSC/internal/cli"
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ayush6624/go-chatgpt"
)

type CreateTest struct {
}

func (c CreateTest) IsMatchCommand() bool {
	moduleName = ""
	flag.StringVar(&moduleName, "test-create", "", "Teste de funcionalidade")
	flag.Parse()
	return moduleName != ""

}
func (c CreateTest) Execute() {
	// obter valor de variavel de ambiente
	var envOpenAiKey = os.Getenv("OPENAI_KEY")
	flag.StringVar(&openAiKey, "openai-key", envOpenAiKey, "Chave de acesso a api do openai")
	flag.StringVar(&pathTestModel, "model-test", "../../test_model.txt", "Teste de funcionalidade")
	flag.StringVar(&codeForTestPath, "code", "", "Caminho do código a ser testado")
	createTest()
}

func (c CreateTest) OnHelp() {
	cli.PrintMessage("Comando para criar testes unitários seguindo um modelo")
}

var pathTestModel string = "../../test_model.txt"

var codeForTestPath string = ""
var openAiKey string = ""

func createTest() {
	if codeForTestPath == "" {
		fmt.Println("Por favor, forneça o caminho do código a ser testado usando a flag -code")
		os.Exit(1)
	}

	var instruct string = "create a test with base in a model: \n" +
		getFile(pathTestModel) + "\n\n" + "for the code: \n" +
		getFile(codeForTestPath)

	generateTest(instruct)
	saveTest(instruct)
}

func generateTest(instruct string) string {
	// generate test with chat gpt 3 api

	client, err := chatgpt.NewClient(openAiKey)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	res, err := client.SimpleSend(ctx, "Hello, how are you?")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	response, err := client.SimpleSend(ctx, instruct)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(res)
	return response.Choices[0].Message.Content + "\n\n" +
		"// Generated by FSC\n" +
		"//"

}

func saveTest(test string) {
	var pathSave = strings.Replace(codeForTestPath, ".dart", "_test.dart", 1)
	pathSave = strings.Replace(pathSave, "/lib", "/test", 1)

	file, err := os.Create(pathSave)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cli.PrintMessage("Teste criado com sucesso")

}

func getFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var text string = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
