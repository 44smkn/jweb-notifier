package inmem

import (
	d "jweb-notifier/domain/user"
)

type UserRepository struct{}

var users = make([]*d.User, 5)

func (u *UserRepository) Register(user *d.User) error {
	users = append(users, user)
	return nil
}

// 後で引数を直す
func (u *UserRepository) Get(userId interface{}) *d.User {
	return nil
}

func (u *UserRepository) ChangeFavorite(user *d.User) error {
	for i, u := range users {
		if user.Id == u.Id {
			users[i] = user
		}
	}
	return nil
}
