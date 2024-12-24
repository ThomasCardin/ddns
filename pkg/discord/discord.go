package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func SendIPChangeNotification(webhookURL string, payload Message) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error encoding discord message: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending webhook message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error sending webhook message %d", resp.StatusCode)
	}

	fmt.Println("discord webhook message sent !")
	return nil
}
