package user

import (
	"errors"
	"regexp"
)

type Password struct {
	val string
}

var charTypeRe = regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`)

func NewPassword(password string) (Password, error) {
	if !charTypeRe.MatchString(password) {
		return Password{}, errors.New("英数字8文字以上ではありません")
	}
	return Password{val: password}, nil
}
