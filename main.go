package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type ChatEntry struct {
	UserInput string `json:"user_input"`
	Response  string `json:"response"`
	TimeStamp string `json:"timestamp"`
}

func main() {
	fmt.Println("Hellof")

	reader := bufio.NewReader(os.Stdin)

	for {
		var input string
		fmt.Print("You: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		response := input
		fmt.Println("bot response: ", response)
		fmt.Println("Timestamp:", time.Now().Format("2006-01-02 15:04:05"))
	}
}
