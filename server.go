package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"
)

type Topic struct {
	Name       string
	Partitions int64
}

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(conn)
	}
}

func handleIncomingRequest(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		command, err := ParseCommand(message, conn)
		if err != nil {
			fmt.Println("error while parsing command :", err)
			continue
		}
		err = command.execute()
		if err != nil {
			fmt.Println("error while executing command:", err)
			continue
		}
		time.Sleep(10 * time.Second)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
