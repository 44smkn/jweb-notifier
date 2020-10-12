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

func (u *UserRepository) Get(userId du.Id) *du.User {
	for i, u := range users {
		if u.GetId() == userId {
			return u
		}
	}
	return nil
}

func (u *UserRepository) ChangeFavorite(user *du.User) error {
	for i, u := range users {
		if user.Equals(u) {
			users[i] = user
		}
	}
	return nil
}
