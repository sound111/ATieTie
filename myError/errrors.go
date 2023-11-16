package myError

import "errors"

var (
	ErrorUserExist        = errors.New("user already exists")
	ErrorServerBusy       = errors.New("server is busy")
	ErrorUserNotExist     = errors.New("user do not exists")
	ErrorPwdInvalid       = errors.New("password is invalid")
	ErrorInvalidID        = errors.New("id is invalid")
	ErrUserNotLogin       = errors.New("user not login")
	ErrPostParamInCorrect = errors.New("url param is incorrect")
)
