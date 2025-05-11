package main

import "fmt"

type ChatEntry struct {
	UserInput string `json:"user_input"`
	Response  string `json:"response"`
	TimeStamp string `json:"timestamp"`
}

func main() {
	fmt.Println("Hellof")
}
