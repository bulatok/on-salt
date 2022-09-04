package salt

import "errors"

var (
	ErrInvalidToken = errors.New("invalid bot token")
	ErrChatNotFound = errors.New("chat not found")
)
