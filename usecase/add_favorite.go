package usecase

import (
	"jweb-notifier/domain/diary"
	dd "jweb-notifier/domain/diary"
	"jweb-notifier/domain/user"
	du "jweb-notifier/domain/user"
	"jweb-notifier/presentation/param"
	"log"

	"github.com/pkg/errors"
)

func init() {
	userId, _ := user.NewId("ttatsuki@gmail.com")
	password, _ := user.NewPassword("67fagagaI0fa")
	user := user.NewUser(userId, password, []diary.Id{})
	err := userRepo.Register(user)
	if err != nil {
		log.Fatalln("テストの初期化に失敗しました")
	}
}

func AddFavorite(f *param.AddFavoriteInput) error {
	userId, err := du.NewId(f.UserId)
	if err != nil {
		return errors.Wrapf(err, "ユーザIDの要件に満たしませんでした. id: %s", userId)
	}
	user := userRepo.Get(userId)
	if user == nil {
		return errors.New("入力したユーザIDに対応するユーザが見つかりませんでした")
	}

	diaryId, err := dd.NewId(f.DiaryId)
	if err != nil {
		return errors.Wrapf(err, "日記IDの要件に満たしませんでした. id: %s", diaryId)
	}

	user.AddFavorite(diaryId)
	return userRepo.ChangeFavorite(user)
}
