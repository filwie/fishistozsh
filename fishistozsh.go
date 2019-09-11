package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

// sed -i -E 's/(- cmd: ): *[0-9]+:0;/\1/' ~/.local/share/fish/fish_history

var cmdPrefix = "- cmd: "
var whenPrefix = "  when: "
var zshHistLineTemplate = ": %s:0;%s"

func getHome() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalln("Could not get current user:", err)
	}
	return user.HomeDir
}

func getFishHistfile() (histfile string) {
	relativePath := "/fish/fish_history"
	if XDGDataHome := os.Getenv("XDG_DATA_HOME"); len(XDGDataHome) > 0 {
		histfile = XDGDataHome + relativePath
	} else {
		histfile = getHome() + "/.local/share/" + relativePath
	}
	if _, err := os.Stat(histfile); os.IsNotExist(err) {
		log.Fatalln("Fish history file does not exist.")
	}
	return histfile
}

func main() {
	file, err := os.Open(getFishHistfile())
	if err != nil {
		log.Fatalln("Could not open Fish history file:", err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		if cmdLine := fileScanner.Text(); strings.HasPrefix(cmdLine, cmdPrefix) {
			cmd := strings.TrimPrefix(cmdLine, cmdPrefix)
			fileScanner.Scan()
			if whenLine := fileScanner.Text(); strings.HasPrefix(whenLine, whenPrefix) {
				when := strings.TrimPrefix(whenLine, whenPrefix)
				zshHistLine := fmt.Sprintf(zshHistLineTemplate, when, cmd)
				fmt.Println(zshHistLine)
			}
		}
	}
}
