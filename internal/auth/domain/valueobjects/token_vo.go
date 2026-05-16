package authvalueobjects

import (
	"errors"
	"strings"
)

var ErrEmptyToken = errors.New(
	"token cannot be empty",
)


type TokenVO struct {
	value string
}

func NewTokenVO(value string) (TokenVO, error) {
	if strings.TrimSpace(value) == "" {
		return TokenVO{}, ErrEmptyToken
	}

	return TokenVO{
		value: value,
	}, nil
}

func (t TokenVO) Value() string {
	return t.value
}