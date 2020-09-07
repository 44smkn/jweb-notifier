package http

import (
	"fmt"

	"github.com/slack-go/slack"
)

type NotifierRepository struct {
	webhookUrl string
	recipient  string
	authorIcon string
}

func (n *NotifierRepository) Notify(text string) error {
	attachment := slack.Attachment{
		Color:      "good",
		AuthorName: "jweb-notifier",
		AuthorIcon: n.authorIcon,
		Text:       fmt.Sprintf("@%s\n%s", n.recipient, text),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(n.webhookUrl, &msg)
	if err != nil {
		return err
	}
	return nil
}
