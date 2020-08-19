package inmem

import "jweb-notifier/domain/model/entity"

type UserRepository struct{}

var users = make([]*entity.User, 5)

func (u *UserRepository) Register(user *entity.User) error {
	users = append(users, user)
	return nil
}
