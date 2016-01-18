# Slack client in Golang

This library allows you to send data into slack group/channel using a Incoming Webhook url.

## How to send data
```go
slack.Send(username, webhookIncoming, text, iconUrl string, attachments ...slackAttachement) error
```

### SlackAttachment struct
```go
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
```
