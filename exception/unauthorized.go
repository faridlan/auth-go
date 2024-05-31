package exception

type UnauthorizedError struct {
	Message string
}

func (n *UnauthorizedError) Error() string {
	return n.Message
}

func NewUnauthorizedError(msg string) error {
	return &UnauthorizedError{
		Message: msg,
	}
}
