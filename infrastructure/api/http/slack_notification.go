package http

import (
	"github.com/slack-go/slack"
)

func NotifyWithSlack(webhookUrl, recipient, text string) error {
	attachment := slack.Attachment{
		Color:      "good",
		AuthorName: "jweb-notifier",
		AuthorIcon: "",
		Text:       text,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(webhookUrl, &msg)
	if err != nil {
		return err
	}
	return nil
}
