package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type CreateTopicCommand struct {
	args map[string]string
}

func (command *CreateTopicCommand) execute() error {
	name := command.args["name"]

	numPartitions, _ := strconv.Atoi(command.args["partitions"])
	numReplicas, _ := strconv.Atoi(command.args["replicas"])

	//create folder
	_, err := os.Stat(name)
	if err == nil {
		fmt.Println("Error: File already exists:", name)
		return fmt.Errorf("error while creating topic: file already exists: %s", name)
	}

	err = os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("error while creating folder")
		return fmt.Errorf("error while creating topic: folder already exists: %s", name)
	}

	for i := 0; i < int(math.Max(float64(numPartitions), 1.0)); i++ {
		for j := 0; j < int(math.Max(float64(numReplicas), 1.0)); j++ {
			fileName := name + "/" + name + strconv.Itoa(i) + "_replica" + strconv.Itoa(j) + ".txt"

			//create file
			_, err = os.Create(fileName)
			if err != nil {
				fmt.Println("error while creating topic:", err)
				return fmt.Errorf("error while creating topic :%s", err)
			}
		}
	}
	return nil
}
