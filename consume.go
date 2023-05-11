package main

import "fmt"

type ConsumeCommand struct {
	args map[string]string
}

func (command *ConsumeCommand) execute() error {
	return fmt.Errorf("error")
}
