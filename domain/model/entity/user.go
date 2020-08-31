package entity

import (
	"jweb-notifier/domain/model/vo"
)

type User struct {
	UserId   *vo.UserId
	Password *vo.Password
}

func NewUser(id *vo.UserId, password *vo.Password) *User {
	return &User{
		UserId:   id,
		Password: password,
	}
}
