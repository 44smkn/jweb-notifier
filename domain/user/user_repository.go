package user

type UserRepository interface {
	Register(*User) error
	ChangeFavorite(*User) error
}
