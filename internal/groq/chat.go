package groq

import (
	"fmt"
	"os"
)

func GetPositiveResponse(userInput string) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	fmt.Println(".env", apiKey)
	return apiKey, nil
}
