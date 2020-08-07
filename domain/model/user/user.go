package user

type User struct {
	UserId   *Id
	Password *Password
}

func New(id *Id, password *Password) *User {
	return &User{
		UserId:   id,
		Password: password,
	}
}
