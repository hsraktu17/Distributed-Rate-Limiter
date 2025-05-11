package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type ChatEntry struct {
	UserInput string `json:"user_input"`
	Response  string `json:"response"`
	TimeStamp string `json:"timestamp"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func main() {
	fmt.Println("Hellof")

	reader := bufio.NewReader(os.Stdin)
	apiKey := os.Getenv("GROQ_API_KEY")
	fmt.Println("Hellof the envs is here:", apiKey)

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
