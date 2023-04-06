package main

import (
    "fmt"
    "math"
    "os"
    "regexp"
    "strconv"
    "strings"
)

type Command interface {
    execute()
}

// implement generics for args
//type Value interface {
//    int64 | string
//}

type CreateTopicCommand struct {
    args map[string]string
}
type ProduceCommand struct {
    args map[string]string
}
type ConsumeCommand struct {
    args map[string]string
}

func (command *CreateTopicCommand) execute() {
    name := command.args["name"]

    numPartitions, _ := strconv.Atoi(command.args["partitions"])
    numReplicas, _ := strconv.Atoi(command.args["replicas"])

    for i := 0; i < int(math.Max(float64(numPartitions), 1.0)); i++ {
        for j := 0; j < int(math.Max(float64(numReplicas), 1.0)); j++ {
            os.Create(name + strconv.Itoa(i) + "_replica" + strconv.Itoa(j))
        }
    }
}
func (command *ProduceCommand) execute() {

}
func (command *ConsumeCommand) execute() {

}

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

func parseCommand(input string) Command {

    splittedInput := strings.Split(input, " ")
    commandType := splittedInput[0]

    var command Command
    switch commandType {
    case "create-topic":
        argsMap, err := getKeyValue(splittedInput[1:])
        if err != nil {
            fmt.Println(err)
            return nil
        }
        for key, value := range argsMap {
            fmt.Println(key + ":" + value)
        }
        createTopic := &CreateTopicCommand{argsMap}
        command = createTopic
    case "produce":
        argsMap := make(map[string]string)
        produceCommand := &ProduceCommand{argsMap}
        command = produceCommand
    case "consume":
        argsMap := make(map[string]string)
        consumeCommand := &ConsumeCommand{argsMap}
        command = consumeCommand
    default:
        fmt.Errorf("unknown command type : %s", commandType)
    }
    return command
}
func main() {
    command := parseCommand("create-topic --name=topicname --partitions=2 --replicas=2")
    if command != nil {
        command.execute()
    }
}
