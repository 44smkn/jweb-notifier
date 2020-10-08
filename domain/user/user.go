package user

import "jweb-notifier/domain/diary"

type User struct {
	id       Id
	password Password
	favorite []diary.Id
}

func NewUser(userId Id, userPassword Password, favorites []diary.Id) *User {
	return &User{
		id:       userId,
		password: userPassword,
		favorite: favorites,
	}
}

func (u *User) AddFavorite(id diary.Id) {
	u.favorite = append(u.favorite, id)
}

func (u *User) Equals(another *User) bool {
	return u.id.val == another.id.val
}
