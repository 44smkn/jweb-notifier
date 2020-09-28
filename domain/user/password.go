package user

import (
	"errors"
	"regexp"
)

type password string

var charTypeRe = regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`)

func NewPassword(userPassword string) (password, error) {
	if !charTypeRe.MatchString(userPassword) {
		return "", errors.New("英数字8文字以上ではありません")
	}
	return password(userPassword), nil
}
