package user

import (
	"jweb-notifier/domain/model/entity"
)

type User interface {
	Register(*entity.User) error
	AddFavorite(*entity.Favorite) error
	DeleteFavorite(*entity.Favorite) error
}
