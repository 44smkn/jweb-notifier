package user

type User struct {
	Id       Id
	Password password
	Favorite []diaryId
}

func NewUser(userId Id, userPassword password, favorites []diaryId) *User {
	return &User{
		Id:       userId,
		Password: userPassword,
		Favorite: favorites,
	}
}

func (u *User) AddFavorite(id diaryId) {
	u.Favorite = append(u.Favorite, id)
}
