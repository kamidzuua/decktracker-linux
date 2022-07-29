package main

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strings"

	"log-proto/utils"

	"github.com/hpcloud/tail"
)

func FindLogDir() string {
	var logdir string
	var prefix string
	var logName string

	user, err := user.Current()
	prefix = "/home/" + user.Username + "/"
	cmd := exec.Command("bash", "-c", "find | grep /Hearthstone/Logs")
	cmd.Dir = prefix
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	logName = "/Power.log"

	dirs := strings.Split(strings.TrimSpace((string(output))), "\n")
	split := strings.Split(dirs[0], "/")

	logdir = strings.Join(split[1:], "/")
	logdir = prefix + logdir + logName
	return logdir
}

//TODO Really Parse Lines
func ParseLine(line string, inBlock *bool) {
	if utils.IsEntity(line) == true {
		utils.ParseEntity(line)
	} else if utils.IsGameStart(line) == true {
		fmt.Println("game started")
	} else if utils.IsGameComplete(line) {
		fmt.Println("game complete")
	} else if utils.IsBlockStart(line) {
		*inBlock = true
	} else if utils.IsBlockEnd(line) {
		*inBlock = false
	} else {
		fmt.Println(line)
	}

	fmt.Println(*inBlock)
}

func main() {

	powloc := FindLogDir()
	inBlock := false
	t, err := tail.TailFile(powloc, tail.Config{Follow: true})

	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		ParseLine(line.Text, &inBlock)
	}
}
