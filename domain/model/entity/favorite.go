package entity

import (
	"jweb-notifier/domain/model/vo"
)

type Favorite struct {
	UserId  *vo.UserId
	DiaryId *vo.DiaryId
}

func NewFavorite(userId *vo.UserId, diaryId *vo.DiaryId) *Favorite {
	return &Favorite{
		UserId:  userId,
		DiaryId: diaryId,
	}
}
