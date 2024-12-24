package discord

import (
	"bytes"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func SendIPChangeNotification(webhookURL string, payload Message) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("error encoding discord message: %w", err)
	}

	log.Infof("message.content: %s", payload.Content)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorf("error sending webhook message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Errorf("error sending webhook message %d", resp.StatusCode)
	}
}
