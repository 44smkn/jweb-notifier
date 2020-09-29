package vo

import (
	"errors"
	"strings"
)

type TalentName string

// リフレクションを利用したほうが良いかも
func NewTalentName(name string) (TalentName, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name is empty")
	}
	return TalentName(name), nil
}
