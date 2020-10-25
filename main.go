package main

//go:generate go-bindata assets/

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	installDictionary()
	fmt.Println()
	fmt.Println("Press the Enter Key to terminate the console screen!")
	_, _ = fmt.Scanln()
}

func installDictionary() {
	// Read Custom Dictionary.txt
	fmt.Println("Retrieving Custom Dictionary.txt ...")
	data, err := Asset("assets/Custom Dictionary.txt")
	if err != nil {
		showErrorMessage("Failed to retrieve Custom Dictionary.txt.\nError message: %s\n", err)
		return
	}
	fmt.Println("Successfully retrieve Custom Dictionary.txt")

	// get local app assets path
	dataPath, err := getAppDataLocalPath()
	if err != nil {
		showErrorMessage("%s", err)
		return
	}

	dictionaryPath := dataPath + "\\Google\\Chrome\\User Data\\Default\\Custom Dictionary.txt"

	fmt.Println("Copying Custom Dictionary to", dictionaryPath)
	err = ioutil.WriteFile(dictionaryPath, data, 0644)
	if err != nil {
		showErrorMessage("Failed to copy Custom Dictionary to "+dictionaryPath+"\nError message: %s", err)
		return
	}

	color.Green("Successfully install anatomy dictionary!!")
}

func getAppDataLocalPath() (string, error) {
	if runtime.GOOS == "windows" {
		localDataPath := os.Getenv("LOCALAPPDATA")
		if localDataPath != "" {
			return localDataPath, nil
		}
		return "", errors.New("Can not find local app assets folder ")
	}
	return "", errors.New("This program can only run on Windows OS ")
}

func showErrorMessage(format string, error error) {
	fmt.Println()
	color.Red(format, error)
}
