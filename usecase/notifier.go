package usecase

import (
	"errors"
	du "jweb-notifier/domain/update"
	"jweb-notifier/infrastructure/notification"
)

const NotificationBySlack = "slack"

type Notifier interface {
	Notify(string) error
}

func NotifyUpdates(upd du.Updates, notificationMethod string) error {
	var notifier Notifier
	switch notificationMethod {
	case NotificationBySlack:
		notifier = notification.NewSlackNotifier()
	default:
		return errors.New("Not exists method.")
	}
	return notifier.Notify()
}
