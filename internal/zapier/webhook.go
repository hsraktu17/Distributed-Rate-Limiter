package zapier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ChatEntry struct {
	UserInput string `json:"user_input"`
	Response  string `json:"response"`
	TimeStamp string `json:"timestamp"`
}

func Send(entry ChatEntry) error {
	webhookURL := os.Getenv("ZAPIER_WEBHOOK_URL")
	body, _ := json.Marshal(entry)

	fmt.Println("Webhook url: ", webhookURL)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
