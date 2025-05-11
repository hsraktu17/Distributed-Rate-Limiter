package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func readPromptFile() string {
	content, err := os.ReadFile(filepath.Join("internal", "groq", "prompt.txt"))
	if err != nil {
		return "you are kind of uplifting chat bot, you main job is to motivate the person talking to you. But motivate them when the person tell you his/her problem"
	}
	return string(content)
}

func GetPositiveResponse(userInput string) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	url := "https://api.groq.com/openai/v1/chat/completions"

	fmt.Println("response: ", readPromptFile())

	reqBody := map[string]any{
		"model": "llama3-70b-8192", // or "mixtral-8x7b-32768"
		"messages": []map[string]string{
			{"role": "system", "content": readPromptFile()},
			{"role": "user", "content": userInput},
		},
	}

	bodyByte, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	var parsed map[string]interface{}
	json.Unmarshal(respBody, &parsed)

	reply := parsed["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return reply, nil
}
