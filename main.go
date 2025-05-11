package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hsraktu17/bot/zapier"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	fmt.Println("Type anything — I’ll respond with something uplifting. 💬")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nYou: ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		userInput := scanner.Text()
		// reply, err := groq.GetPositiveResponse(userInput)
		reply := "hello"

		if err != nil {
			fmt.Println("Bot: (there is some bot error)")
			fmt.Println("Groq Error", err)
			continue
		}

		fmt.Println(reply)

		entry := zapier.ChatEntry{
			UserInput: userInput,
			Response:  reply,
			TimeStamp: time.Now().Format(time.RFC3339),
		}

		err = zapier.Send(entry)

		if err != nil {
			log.Println("Failed to send to zapier", err)
		}
	}
}
