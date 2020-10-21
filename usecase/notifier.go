package usecase

import (
	"errors"
	"fmt"
	du "jweb-notifier/domain/update"
	"jweb-notifier/infrastructure/notification"
	"time"
)

const NotificationBySlack = "slack"

type Notifier interface {
	Notify(string) error
}

func NotifyUpdates(upd du.Updates, notificationMethod string) error {
	var notifier Notifier
	switch notificationMethod {
	case NotificationBySlack:
		// TODO: 設定からURLを注入できるようにするか取得処理をいれる
		notifier = notification.NewSlackNotifier("webhook")
	default:
		return errors.New("Not exists method.")
	}
	header := `更新日時： %v
あなたのお気に入りのブログに更新がありました！`
	return notifier.Notify(fmt.Sprintf(header, upd.GetDateTime(time.RFC822)))
}
