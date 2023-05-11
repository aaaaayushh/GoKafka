package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type ProduceCommand struct {
	args map[string]string
	conn net.Conn
}

func (command *ProduceCommand) execute() error {
	topicName := command.args["name"]
	pathToFile := topicName + "/" + topicName + "0_replica0.txt"
	file, err := os.OpenFile(pathToFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return fmt.Errorf("error while opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(command.conn)

	for scanner.Scan() {
		message := scanner.Text()
		byteMessage := []byte(message)
		if err != nil {
			fmt.Println("error reading message", err)
			return fmt.Errorf("error reading message")
		}
		_, err = file.Write(byteMessage)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return nil
}
