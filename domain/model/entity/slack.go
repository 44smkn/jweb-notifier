package entity

import (
	"jweb-notifier/domain/model/vo"
)

type Slack struct {
	UserId     vo.UserId
	WebhookUrl string
	Recipient  string
	Token      string
}

func (s *Slack) Notify(msg string) error {
	return nil
}
