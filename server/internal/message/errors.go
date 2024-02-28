package message

import "errors"

var (
	ErrUnableToVerify            = errors.New("unable to verify if request was from a bot")
	ErrIsBot                     = errors.New("request came from a bot")
	ErrUnableToInsertMessage     = errors.New("unable to save message")
	ErrCouldNotReadBody          = errors.New("request body is malformed")
	ErrInvalidMessageRequestBody = errors.New("request body is not a correct message body")
)
