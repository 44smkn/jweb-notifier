package update

import (
	d "jweb-notifier/domain/diary"
)

type Updates struct {
	Date    string
	Time    string
	Diaries []d.Diary
}
