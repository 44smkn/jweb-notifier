package user

import . "jweb-notifier/domain/favorite"

type UserRepository interface {
	Register(*User) error
	AddFavorite(*Favorite) error
	DeleteFavorite(*Favorite) error
}
