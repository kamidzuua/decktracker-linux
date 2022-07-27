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
	var power string

	user, err := user.Current()
	prefix = "/home/" + user.Username + "/"
	cmd := exec.Command("bash", "-c", "find | grep /Hearthstone/Logs")
	cmd.Dir = prefix
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	power = "/Power.log"

	dirs := strings.Split(strings.TrimSpace((string(output))), "\n")
	split := strings.Split(dirs[0], "/")

	logdir = strings.Join(split[1:], "/")
	logdir = prefix + logdir + power
	return logdir
}

//TODO Really Parse Lines
func ParseLine(line string) {
	if utils.IsEntity(line) == true {
		utils.ParseEntity(line)
	} else if utils.IsGameStart(line) == true {
		fmt.Println("game started")
	}
}

func main() {

	powloc := FindLogDir()

	t, err := tail.TailFile(powloc, tail.Config{Follow: true})

	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		ParseLine(line.Text)
	}
}
