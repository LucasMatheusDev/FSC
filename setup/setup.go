package main

import (
	"FSC/internal/cli"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	cli.PrintMessage("Starting setup...")

	switch runtime.GOOS {
	case "windows":
		setupWindows()
	case "linux", "darwin":
		setupUnix()
	default:
		cli.PrintMessage("Setup is not supported on this operating system.")
		os.Exit(1)
	}
}

var pathMainExecutable string = "./main.exe"
var executableName string = "fsc.exe"
var fileName string = "fsc"

func setupWindows() {
	cli.PrintMessage("Setting up for Windows...")
	userPath, err := os.UserHomeDir()
	if err != nil {
		cli.PrintMessage("Error getting user home directory: " + err.Error())
		return
	}

	fmt.Println("User path: " + userPath)

	err = copyFile(pathMainExecutable, filepath.Join(userPath, "\\"+executableName))
	if err != nil {
		fmt.Println("Error moving "+executableName+": ", err)
		return
	}

	err = addToPath(filepath.Join(userPath, fileName), true)
	if err != nil {
		fmt.Println("Error adding "+executableName+" to PATH: ", err)
		return
	}

	cli.PrintMessage("Setup completed successfully.")
}

func setupUnix() {
	cli.PrintMessage("Setting up for Unix...")

	userPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory: ", err)
		return
	}

	err = copyFile(executableName, filepath.Join(userPath, executableName))

	if err != nil {
		cli.PrintMessage("Error moving" + executableName + ": " + err.Error())
		return
	} else {
		addToPath(filepath.Join(userPath, "fsc"), false)
	}
}

func copyFile(source, destination string) error {
	fmt.Println("Copying " + source + " to " + destination)

	if _, err := os.Stat(destination); err == nil {
		if err := os.Remove(destination); err != nil {
			return fmt.Errorf("failed to remove existing destination file: %w", err)
		}
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to get source file on: %w "+source, err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w "+" Destination: "+destination, err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file to destination %w"+" Destination: "+destination, err)
	}

	return nil
}

func addToPath(path string, isWindows bool) error {
	fmt.Println("Adding to PATH: " + path)

	var err error = nil
	if isWindows {
		err = exec.Command("setx", "PATH", `%PATH%;`+path).Run()

	} else {
		err = exec.Command("export", "PATH", `$PATH:`+path).Run()
	}

	if err != nil {
		return fmt.Errorf("failed to add to PATH: %w", err)
	}
	cli.PrintMessage("Added to PATH: " + path)
	return err
}
