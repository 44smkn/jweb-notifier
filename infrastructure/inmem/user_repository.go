package inmem

import "jweb-notifier/domain/model/user"

type UserRepository struct{}

func (u *UserRepository) Register(user *user.User) error {
	return nil
}
