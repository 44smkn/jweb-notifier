package usecase

import (
	"jweb-notifier/domain/user"
	"jweb-notifier/presentation/param"

	"github.com/pkg/errors"
)

func AddFavorite(f *param.Favorite) error {
	userId, err := user.NewId(f.UserId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userId is fail. id: %s", f.UserId)
	}
	diaryId, err := user.NewDiaryId(f.DiaryId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of diaryId is fail. id: %s", f.DiaryId)
	}
	return repo.AddFavorite(fav)
}
