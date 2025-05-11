package groq

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func GetPositiveResponse(userInput string) ([]string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	url := " https://api.groq.com/openai/v1/chat/completions"

	reqBody := map[string]interface{}{
		"model": "llama-3.3-70b-versatile",
		"message": []map[string]string{
			{"role": "system", "content": "You are a kind uplifting chatbot, always respond with something positive and the person talking to you is going through ups and downs in his life"},
			{"role": "user", "content": userInput},
		},
	}

	bodyByte, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	req.Header.Set("Content-Type", "application/json")
	return []string{apiKey, url}, nil
}
