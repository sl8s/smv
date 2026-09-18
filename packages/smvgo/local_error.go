package smvgo

import (
	"fmt"
)

type LocalError interface {
	error
	EnumGuilty() EnumGuilty
	Message() string
	final()
}

type localError struct {
	source     string
	enumGuilty EnumGuilty
	message    string
}

func (le localError) Error() string {
	return fmt.Sprintf("LocalError{source: %s, enumGuilty: %s, message: %s}", le.source, le.enumGuilty.String(), le.message)
}

func (le localError) EnumGuilty() EnumGuilty {
	return le.enumGuilty
}

func (le localError) Message() string {
	return le.message
}

func (le localError) final() {
}

func NewLocalError(source string, enumGuilty EnumGuilty, message string) LocalError {
	return localError{
		source:     source,
		enumGuilty: enumGuilty,
		message:    message,
	}
}
