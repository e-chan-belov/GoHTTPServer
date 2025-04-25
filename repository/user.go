package repository

import "hw2/domain"

type User interface {
	Post(user domain.User) error
	Put(user domain.User) error
	Get(id domain.UserId) (domain.User, error)
}

type Session interface {
	Post(session domain.Session) error
	Put(session domain.Session) error
	Get(id domain.SessionId) (domain.Session, error)
	Delete(id domain.SessionId) error
}
