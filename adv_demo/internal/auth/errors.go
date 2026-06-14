package auth

import "errors"

var (
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrInvalidCredentials = errors.New("invalid credentials")
) 
