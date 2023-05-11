package main

import (
	"fmt"
	"net"
	"strings"
)

type Command interface {
	execute() error
}

// implement generics for args
//type Value interface {
//    int64 | string
//}

func ParseCommand(input string, conn net.Conn) (Command, error) {

	splitInput := strings.Split(input, " ")
	commandType := splitInput[0]

	var command Command
	var err error
	switch commandType {
	case "create-topic":
		command, err = handleCommand(splitInput[1:], commandType, conn)
		if err != nil {
			fmt.Println("error in create topic:", err)
			return nil, err
		}
	case "produce":
		command, err = handleCommand(splitInput[1:], commandType, conn)
		if err != nil {
			fmt.Println("error in producing message:", err)
			return nil, err
		}
	case "consume":
		command, err = handleCommand(splitInput[1:], commandType, conn)
	default:
		return nil, fmt.Errorf("unknown command type : %s", commandType)
	}
	return command, nil
}
