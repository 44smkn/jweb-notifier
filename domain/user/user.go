package user

type User struct {
	id       Id
	password password
	favorite []diaryId
}

func NewUser(userId Id, userPassword password, favorites []diaryId) *User {
	return &User{
		id:       userId,
		password: userPassword,
		favorite: favorites,
	}
}

func (u *User) AddFavorite(id diaryId) {
	u.favorite = append(u.favorite, id)
}

func (u *User) Equals(another *User) bool {
	return u.id.val == another.id.val
}
