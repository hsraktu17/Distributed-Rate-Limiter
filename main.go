package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/hsraktu17/bot/internal/groq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nYou: ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		userInput := scanner.Text()
		reply, err := groq.GetPositiveResponse(userInput)

		if err != nil {
			fmt.Println("Error", err)
		}

		fmt.Println(reply)
	}
}
