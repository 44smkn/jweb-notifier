package usecase

import (
	"jweb-notifier/infrastructure/persistence/inmem"
	d "jweb-notifier/infrastructure/persistence/inmem"
)

var userRepo *d.UserRepository

func init() {
	userRepo = &inmem.UserRepository{}
}
