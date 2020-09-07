package vo

import (
	"errors"
	"regexp"
)

type UserId string

// mailRe はメールアドレスの正規表現
var mailRe = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func NewUserId(id string) (UserId, error) {
	if !mailRe.MatchString(id) {
		return "", errors.New("IDがメールアドレスの形式になっていません")
	}
	return UserId(id), nil
}
