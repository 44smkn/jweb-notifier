package usecase

import (
	"jweb-notifier/domain/model/repository"
	"jweb-notifier/infrastructure/inmem"
)

var repo repository.User

func init() {
	repo = &inmem.UserRepository{}
}
