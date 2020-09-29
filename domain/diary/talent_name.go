package diary

import (
	"errors"
	"strings"
)

type talentName string

// リフレクションを利用したほうが良いかも
func NewTalentName(name string) (talentName, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name is empty")
	}
	return talentName(name), nil
}
