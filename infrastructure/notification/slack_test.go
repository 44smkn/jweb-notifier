package notification_test

import (
	"fmt"
	"jweb-notifier/infrastructure/notification"
	"os"
	"testing"
)

func TestSlackNotify(t *testing.T) {
	webhookurl := os.Getenv("SLACK_WEBHOOK_URL")
	fmt.Println(webhookurl)
	n := notification.NewSlackNotifier(webhookurl)
	err := n.Notify("Test Message")
	if err != nil {
		t.Errorf("failed to notify:  %v", err)
	}
}
