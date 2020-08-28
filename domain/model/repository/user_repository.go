package user

import (
	"jweb-notifier/domain/model/entity"
)

type Repository interface {
	Register(*entity.User) error
}
