package inmem

import (
	. "jweb-notifier/domain/user"
)

type UserRepository struct{}

var users = make([]*User, 5)

func (u *UserRepository) Register(user *User) error {
	users = append(users, user)
	return nil
}

func (u *UserRepository) ChangeFavorite(user *User) error {
	for i, u := range users {
		if user.Id == u.Id {
			users[i] = user
		}
	}
	return nil
}
