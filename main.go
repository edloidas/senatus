package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=Ping Pong=")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		run(command)
	}
}

func run(command string) {
	switch command {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "ping":
		fmt.Println("pong")
		break
	case "pong":
		fmt.Println("ping")
		break
	default:
		fmt.Printf("Unknown command '%s'.\n", command)
	}
}
