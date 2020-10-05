package usecase

import (
	dd "jweb-notifier/domain/diary"
	du "jweb-notifier/domain/user"
	"jweb-notifier/presentation/param"

	"github.com/pkg/errors"
)

func RegisterUser(u *param.User) error {
	id, err := du.NewId(u.Id)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userid is fail. id: %s", u.Id)
	}
	password, err := du.NewPassword(u.Password)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userid is fail. password: %s", u.Password)
	}
	user := du.NewUser(id, password, []dd.Id{})
	return userRepo.Register(user)
}
