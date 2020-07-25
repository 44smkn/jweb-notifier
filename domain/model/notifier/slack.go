package notifier

import "jweb-notifier/domain/model/user"

type Slack struct {
	UserId     user.Id
	WebhookUrl string
	Recipient  string
	Token      string
}

func (s *Slack) Notify(msg string) error {
	return nil
}
