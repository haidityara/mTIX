package constant

import "errors"

var ErrorInvalidLogin = errors.New("invalid email or password")
var ErrorEmailAlreadyExists = errors.New("email already exists")
var ErrorInvalidRole = errors.New("invalid role")
