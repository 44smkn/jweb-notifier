package user

type Repository interface {
	Register(*User) error
}
