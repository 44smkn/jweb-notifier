package vo

import (
	"errors"
	"regexp"
)

type UserId struct {
	id string
}

// mailRe はメールアドレスの正規表現
var mailRe = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func NewUserId(id string) (*UserId, error) {
	if !mailRe.MatchString(id) {
		return nil, errors.New("IDがメールアドレスの形式になっていません")
	}
	return &UserId{id}, nil
}

func (u *UserId) Equals(another *UserId) bool {
	return u.id == another.id
}
