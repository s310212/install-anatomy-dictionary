package main

//go:generate go-bindata assets/

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/sqweek/dialog"
	"io/ioutil"
	"os/exec"
)

func main() {
	stopChrome()
	installDictionary()
	pause()
}

func stopChrome() {
	if isProcessRunning("chrome.exe") {
		ok := dialog.Message("%s", "Install program can run only when Chrome is not running.\nAre you sure to kill Chrome process?").Title("Kill Chrome process?").YesNo()
		if !ok {
			showErrorAndExit("%s", errors.New("Install program can not run when Chrome is running."))
		} else {
			_, err := exec.Command("taskkill", "/F", "/IM", "chrome.exe").Output()
			if err != nil {
				showErrorAndExit("%s", errors.New("Cannot stop Chrome. Please contact the developer."))
			}
		}
	}
}

func installDictionary() {
	// Read Custom Dictionary.txt
	fmt.Println("Retrieving Custom Dictionary.txt ...")
	data, err := Asset("assets/Custom Dictionary.txt")
	if err != nil {
		showErrorAndExit("Failed to retrieve Custom Dictionary.txt.\nError message: %s\n", err)
	}
	fmt.Println("Successfully retrieve Custom Dictionary.txt")

	// get local app assets path
	dataPath, err := getAppDataLocalPath()
	if err != nil {
		showErrorAndExit("%s", err)
	}

	dictionaryPath := dataPath + "\\Google\\Chrome\\User Data\\Default\\Custom Dictionary.txt"

	fmt.Println("Copying Custom Dictionary to", dictionaryPath)
	err = ioutil.WriteFile(dictionaryPath, data, 0644)
	if err != nil {
		showErrorAndExit("Failed to copy Custom Dictionary to "+dictionaryPath+"\nError message: %s", err)
	}

	color.Green("Successfully install anatomy dictionary!!")
}
