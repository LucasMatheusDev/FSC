package utils

import (
	"FSC/internal/cli"
	"os"
	"os/exec"
	"runtime"
)

func GetProjectPath() string { return os.Getenv("FS_PROJECT_PATH") }

func SetProjectPath(projectPath string) (string, error) {
	var isWindows bool = runtime.GOOS == "windows"
	var err error = nil
	if isWindows {
		err = exec.Command("setx", "FS_PROJECT_PATH", projectPath).Run()

	} else {
		err = exec.Command("export", "FS_PROJECT_PATH", projectPath).Run()
	}

	if err != nil {
		cli.PrintMessage("failed to add to PATH: " + err.Error())
		return "", err
	} else {
		return projectPath, nil
	}

}
