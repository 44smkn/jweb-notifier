package user

import (
	"errors"
	"regexp"
)

type id string

// mailRe はメールアドレスの正規表現
var mailRe = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func NewId(userId string) (id, error) {
	if !mailRe.MatchString(userId) {
		return "", errors.New("IDがメールアドレスの形式になっていません")
	}
	return id(userId), nil
}
