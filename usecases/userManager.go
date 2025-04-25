package usecases

import "hw2/domain"

type UserManager interface {
	Register(username domain.UserLogin, pass domain.UserPassword) error
	Login(username domain.UserLogin, pass domain.UserPassword) (domain.SessionId, error)
	Auth(id domain.SessionId) error
}
