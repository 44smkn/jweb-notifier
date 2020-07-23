package user

import (
	"errors"
	"regexp"
)

type Id struct {
	id string
}

// mailRe はメールアドレスの正規表現
var mailRe = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func NewId(id string) (*Id, error) {
	if !mailRe.MatchString(id) {
		return nil, errors.New("IDがメールアドレスの形式になっていません")
	}
	return &Id{id}, nil
}
