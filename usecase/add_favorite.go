package usecase

import (
	dd "jweb-notifier/domain/diary"
	du "jweb-notifier/domain/user"

	"github.com/pkg/errors"
)

func AddFavorite(userId, diaryId string) error {
	uid, err := du.NewId(userId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userId is fail. id: %s", userId)
	}
	user := userRepo.Get(uid)

	did, err := dd.NewId(diaryId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of diaryId is fail. id: %s", diaryId)
	}

	user.AddFavorite(did)
	return userRepo.ChangeFavorite(user)
}
