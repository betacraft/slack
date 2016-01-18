package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type slackMessage struct {
	Username    string            `json:"username"`
	Text        string            `json:"text"`
	IconUrl     string            `json:"icon_url"`
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Fallback string `json:"fallback"`
	Color    string `json:"color"`
	Pretext  string `json:"pretext"`

	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`

	Title     string `json:"title"`
	TitleLink string `json:"title_link"`

	Text string `json:"text"`

	MarkdownIn string `json:"mrkdwn_in"`

	Fields []SlackAttachmentField `json:"fields"`
}

type SlackAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func Send(username, webhookIncoming, text, iconUrl string, attachments ...SlackAttachment) error {
	message := slackMessage{Username: username, Text: text, IconUrl: iconUrl, Attachments: attachments}
	return send(webhookIncoming, &message)
}

func send(webhookIncoming string, message *slackMessage) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	resp, err := http.Post(webhookIncoming, "application/json", bytes.NewBuffer(jsonMessage))
	log.Println(resp)
	return err
}
