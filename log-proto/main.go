package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var powloc = "/home/kamidzu/.wine.hearthstone/dosdevices/c:/Program Files/Hearthstone/Logs/Power.log"

func ParseLine(line string) {
	fmt.Println(line)
}

func main() {
	t, err := tail.TailFile(powloc, tail.Config{Follow: true})

	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		ParseLine(line.Text)
	}
}
