package repository

import (
	"jweb-notifier/domain/model/entity"
)

type User interface {
	Register(*entity.User) error
}
