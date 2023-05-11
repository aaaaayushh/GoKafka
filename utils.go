package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

func validateKeyValue(input string) bool {
	re := regexp.MustCompile(`--\w+=\w+`)
	return re.MatchString(input)
}

func getKeyValue(input []string) (map[string]string, error) {
	argsMap := make(map[string]string)

	for _, substr := range input {
		if validateKeyValue(substr) {
			keyValue := strings.Split(substr, "=")
			argsMap[keyValue[0][2:]] = keyValue[1]
		} else {
			return nil, fmt.Errorf("incorrect syntax for key value:%s", substr)
		}
	}

	return argsMap, nil
}

// include this inside command struct ??
func handleCommand(input []string, commandType string, conn net.Conn) (Command, error) {
	var command Command
	argsMap, err := getKeyValue(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for key, value := range argsMap {
		fmt.Println(key + ":" + value)
	}
	switch commandType {
	case "create-topic":
		command = &CreateTopicCommand{argsMap}
	case "produce":
		command = &ProduceCommand{argsMap, conn}
	case "consume":
		command = &ConsumeCommand{argsMap}
	}
	return command, nil
}
