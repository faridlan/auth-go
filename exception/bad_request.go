package exception

import "strings"

type BadRequestError struct {
	Msg []string
}

func (b *BadRequestError) Error() string {
	return strings.Join(b.Msg, ",")
}

func NewBadRequestError(msg []string) error {
	return &BadRequestError{
		Msg: msg,
	}
}
