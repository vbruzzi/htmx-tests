package errors

import "net/http"

const (
	EINVALID  = "invalid"
	EINTERNAL = "internal"
)

type Error struct {
	Code   string
	Status int
	Err    error
}

func NewError(err error, code string) *Error {
	return &Error{
		Err:    err,
		Code:   code,
		Status: GetResponseFromCode(code),
	}
}

func GetResponseFromCode(code string) int {
	if code == EINVALID {
		return http.StatusUnprocessableEntity
	}
	if code == EINTERNAL {
		return http.StatusInternalServerError
	}
	panic("invalid_code")
}
