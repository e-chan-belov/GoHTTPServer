package repository

import "errors"

var AlreadyExists error = errors.New("DB error! Already exists")
var NotFound error = errors.New("DB error! not found")
var WrongToken error = errors.New("DB error! wrong token")
