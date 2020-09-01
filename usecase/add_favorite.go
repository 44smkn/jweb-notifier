package usecase

import (
	"jweb-notifier/domain/model/entity"
	"jweb-notifier/domain/model/vo"
	"jweb-notifier/presentation/param"

	"github.com/pkg/errors"
)

func AddFavorite(f *param.Favorite) error {
	userId, err := vo.NewUserId(f.UserId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userId is fail. id: %s", f.UserId)
	}
	diaryId, err := vo.NewDiaryId(f.DiaryId)
	if err != nil {
		return errors.Wrapf(err, "creating vo of diaryId is fail. id: %s", f.DiaryId)
	}
	fav := entity.NewFavorite(userId, diaryId)
	return repo.AddFavorite(fav)
}
