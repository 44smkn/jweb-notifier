package usecase

import (
	dd "jweb-notifier/domain/diary"
	du "jweb-notifier/domain/user"
	"jweb-notifier/presentation/param"

	"github.com/pkg/errors"
)

func AddFavorite(f *param.Favorite) error {
	userId, err := du.NewId(f.UserId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userId is fail. id: %s", f.UserId)
	}
	user := userRepo.Get(userId)

	diaryId, err := dd.NewId(f.DiaryId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of diaryId is fail. id: %s", f.DiaryId)
	}

	user.AddFavorite(diaryId)
	return userRepo.ChangeFavorite(user)
}
