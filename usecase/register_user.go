package usecase

import (
	"jweb-notifier/domain/model/user"
	"jweb-notifier/infrastructure/inmem"
	"jweb-notifier/presentation/param"

	"github.com/pkg/errors"
)

var repo user.Repository = &inmem.UserRepository{}

func RegisterUser(u *param.User) error {
	id, err := user.NewId(u.Id)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userid is fail. id: %s", u.Id)
	}
	password, err := user.NewPassword(u.Password)
	if err != nil {
		return errors.Wrapf(err, "creating vo of userid is fail. password: %s", u.Password)
	}
	user := user.New(id, password)
	return repo.Register(user)
}
