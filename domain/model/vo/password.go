package vo

import (
	"errors"
	"regexp"
)

type Password struct {
	password string
}

var charTypeRe = regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`)

func NewPassword(password string) (*Password, error) {
	if !charTypeRe.MatchString(password) {
		return nil, errors.New("英数字8文字以上ではありません")
	}
	return &Password{password}, nil
}
