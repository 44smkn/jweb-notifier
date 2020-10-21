package update

import (
	d "jweb-notifier/domain/diary"
)

type Updates struct {
	dateTime DateTime
	diaries  []d.Diary
}

func NewUpdates(dateTime DateTime, diaries []d.Diary) *Updates {
	return &Updates{
		dateTime: dateTime,
		diaries:  diaries,
	}
}

func (u *Updates) GetDateTime(layout string) string {
	return u.dateTime.Format(layout)
}
