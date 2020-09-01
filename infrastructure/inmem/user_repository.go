package inmem

import "jweb-notifier/domain/model/entity"

type UserRepository struct{}

var users = make([]*entity.User, 5)
var favs = make([]*entity.Favorite, 5)

func (u *UserRepository) Register(user *entity.User) error {
	users = append(users, user)
	return nil
}

func (u *UserRepository) AddFavorite(fav *entity.Favorite) error {
	favs = append(favs, fav)
	return nil
}
