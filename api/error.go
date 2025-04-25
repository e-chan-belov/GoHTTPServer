package api

import "errors"

var NoId error = errors.New("no id")
var NoAuth error = errors.New("no token")
