package usecase

import (
	dd "jweb-notifier/domain/diary"
	du "jweb-notifier/domain/user"
	"jweb-notifier/presentation"

	"github.com/pkg/errors"
)

func RegisterUser(u *presentation.RegisterUserParam) error {
	id, err := du.NewId(u.Id)
	if err != nil {
		return errors.Wrapf(err, "ユーザIDは要件に満たしませんでした. id: %s", u.Id)
	}
	password, err := du.NewPassword(u.Password)
	if err != nil {
		return errors.Wrapf(err, "パスワードは要件に満たしませんでした. password: %s", u.Password)
	}
	user := du.NewUser(id, password, []dd.Id{})
	return userRepo.Register(user)
}
