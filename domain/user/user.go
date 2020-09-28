package user

type User struct {
	Id       id
	Password password
}

func NewUser(userId id, userPassword password) *User {
	return &User{
		Id:       userId,
		Password: userPassword,
	}
}
