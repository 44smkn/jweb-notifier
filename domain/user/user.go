package user

type User struct {
	Id       id
	Password password
	Favorite []diaryId
}

func NewUser(userId id, userPassword password, favorites []diaryId) *User {
	return &User{
		Id:       userId,
		Password: userPassword,
		Favorite: favorites,
	}
}
