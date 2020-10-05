package diary

import (
	"errors"
	"strings"
)

type TalentName struct {
	val string
}

// リフレクションを利用したほうが良いかも
func NewTalentName(name string) (TalentName, error) {
	if strings.TrimSpace(name) == "" {
		return TalentName{}, errors.New("name is empty")
	}
	return TalentName{val: name}, nil
}
