package favorite

import (
	"jweb-notifier/domain/model/diary"
	"jweb-notifier/domain/model/user"
)

type Favorite struct {
	UserId  user.Id
	DiaryId diary.Id
}
