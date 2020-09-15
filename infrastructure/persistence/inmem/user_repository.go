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

func (u *UserRepository) DeleteFavorite(fav *entity.Favorite) error {
	for i, v := range favs {
		if v.UserId == fav.UserId && v.DiaryId == fav.DiaryId {
			favs[i] = favs[len(favs)-1] // Copy last element to index i.
			favs[len(favs)-1] = nil     // Erase last element (write zero value).
			favs = favs[:len(favs)-1]   // Truncate slice.
		}
	}
	return nil
}
