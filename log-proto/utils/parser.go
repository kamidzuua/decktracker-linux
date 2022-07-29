package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type EntityObject struct {
	entityName string
	id         int
	zone       string
	zonePos    int
	cardId     string
	player     int
}

type Block struct {
	startString string
	content     []string
	endString   string
}

func GameMode(line string) string {
	return "not implemented yet"
}

func IsEntity(line string) bool {
	return strings.Contains(line, "entityName=")
}

func IsBlockStart(line string) bool {
	return strings.Contains(line, "BLOCK_START")
}

func IsBlockEnd(line string) bool {
	return strings.Contains(line, "BLOCK_END")
}

func IsGameStart(line string) bool {
	return strings.Contains(line, "CREATE_GAME")
}

func IsGameComplete(line string) bool {
	return strings.Contains(line, "value=COMPLETE")
}

func ParseEntity(line string) {
	fmt.Println(makeEntStruct(line))
}

func makeEntStruct(line string) EntityObject {
	split := strings.Split(strings.Split(line, "[")[1], " ")
	fmt.Println(split)
	result := EntityObject{}
	var parsed []string
	var res []string

	for i := 0; i < len(split); i++ {
		parsed = nil
		parsed = append(parsed, split[i])
		if strings.Contains(split[i], "]") {
			word := strings.Split(split[i], "]")[0]
			res = append(res, word)
			break
		}
		if i+1 < len(split) && !strings.Contains(split[i+1], "=") {
			parsed = append(parsed, split[i+1])
			res = append(res, strings.Join(parsed, " "))
			continue
		}

		res = append(res, split[i])

	}

	for _, value := range res {
		var err error
		curr := strings.Split(value, "=")


		switch curr[0] {
		case "entityName":
			result.entityName = curr[1]
		case "cardId":
			result.cardId = curr[1]
		case "id":
			result.id, err = strconv.Atoi(curr[1])
		case "player":
			result.player, err = strconv.Atoi(curr[1])
		case "zone:
			result.zone = curr[1]
		case "zonePos":
			result.zonePos, err = strconv.Atoi(curr[1])
		}

		if err != nil {
			fmt.Println(err)
		}
	}

	return result
}
