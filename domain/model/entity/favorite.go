package entity

import (
	"jweb-notifier/domain/model/vo"
)

type Favorite struct {
	UserId  vo.UserId
	DiaryId vo.DiaryId
}
