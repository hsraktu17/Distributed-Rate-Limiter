package groq

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func GetPositiveResponse(userInput string) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	url := "https://api.groq.com/openai/v1/chat/completions"

	reqBody := map[string]any{
		"model": "llama3-70b-8192", // or "mixtral-8x7b-32768"
		"messages": []map[string]string{
			{"role": "system", "content": "You are a kind uplifting chatbot. Always respond with something positive. The user is going through ups and downs in life. Reply normally until the user tell you his problem and after that you have to work like his therapiest... Reply in mockery way"},
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
