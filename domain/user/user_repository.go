package user

type UserRepository interface {
	Register(*User) error
	AddFavorite(*User) *User
	DeleteFavorite(*User) error
}
