package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func isProcessRunning(name string) bool {
	out, err := exec.Command("TASKLIST", "/FI", "IMAGENAME eq " + name, "/FO", "CSV", "/NH").Output()
	if err != nil {
		return true
	} else {
		outStr := string(out)
		if strings.HasPrefix(outStr, `"chrome.exe"`) {
			return true
		}
		return false
	}
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

func pause() {
	fmt.Println()
	fmt.Println("Press the Enter Key to continue...")
	_, _ = fmt.Scanln()
}

func showErrorAndExit(format string, error error) {
	fmt.Println()
	color.Red("Error: " + format, error)
	pause()
	panic("")
}
