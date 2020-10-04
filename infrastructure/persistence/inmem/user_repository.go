package inmem

import (
	du "jweb-notifier/domain/user"
)

type UserRepository struct{}

var users = make([]*du.User, 5)

func (u *UserRepository) Register(user *du.User) error {
	users = append(users, user)
	return nil
}

// 後で引数を直す
func (u *UserRepository) Get(userId du.Id) *du.User {
	return nil
}

func (u *UserRepository) ChangeFavorite(user *du.User) error {
	for i, u := range users {
		if user.Id == u.Id {
			users[i] = user
		}
	}
	return nil
}
