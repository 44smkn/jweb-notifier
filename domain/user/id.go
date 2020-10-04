package user

import (
	"errors"
	"regexp"
)

// TODO: この状態だと外部パッケージから構造体リテラルを利用して初期化できてしまうため、linter使って防ぐ
type Id struct {
	val string
}

// mailRe はメールアドレスの正規表現
var mailRe = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func NewId(userId string) (Id, error) {
	if !mailRe.MatchString(userId) {
		return Id{}, errors.New("IDがメールアドレスの形式になっていません")
	}
	return Id{val: userId}, nil
}
