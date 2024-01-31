package utils

import (
	"FSC/internal/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetProjectPath() string {

	baseFolder := BaseFolder()
	cli.PrintVerboseMessage("Base folder: " + baseFolder)
	if baseFolder == "" {
		return ""
	}

	filePath := filepath.Join(baseFolder, "project_path.txt")
	cli.PrintVerboseMessage("File path: " + filePath)

	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		cli.PrintMessage("Error opening file: " + err.Error())
		return ""
	}

	file, err = os.Open(filePath)
	if err != nil {
		cli.PrintMessage("failed to open file: %w" + err.Error())
		return ""
	}
	defer file.Close()

	// Use ioutil.ReadAll para ler todo o conteúdo do arquivo para []byte
	data, err := ioutil.ReadAll(file)
	if err != nil {
		cli.PrintMessage("failed to read file: %w" + err.Error())
		return ""
	}

	content := string(data)

	cli.PrintVerboseMessage("Project Path: " + content)

	return content

}

func SetProjectPath(projectPath string) (string, error) {
	
	cli.PrintVerboseMessage("Project Path: " + projectPath)
	if strings.Contains(projectPath, "\"") || strings.Contains(projectPath, "'") {
		cli.PrintVerboseMessage("Project Path contains quotes, please remove them")
		os.Exit(1)
		return "", nil
	}

	baseFolder := BaseFolder()
	if baseFolder == "" {
		return "", nil
	}

	filePath := filepath.Join(baseFolder, "project_path.txt")

	os.MkdirAll(baseFolder, os.ModePerm)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", err
	}

	_, err = file.WriteString(projectPath)
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	cli.PrintVerboseMessage("Project Path saved successfully")
	cli.PrintVerboseMessage("Project Path Saved: " + projectPath)

	return projectPath, nil
}

func BaseFolder() string {
	userPath, err := os.UserHomeDir()
	if err != nil {
		cli.PrintMessage("Error getting user home directory: " + err.Error())
		return ""
	} else {
		return filepath.Join(userPath, ".fsc/")
	}
}
