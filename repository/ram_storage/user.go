package ram_storage

import (
	"hw2/domain"
	"hw2/repository"
)

type UserLine struct {
	Username        domain.UserLogin
	Password        domain.UserPassword
	ActiveSessionId domain.SessionId
}

type User struct {
	data map[domain.UserId]UserLine
}

func userToLine(user domain.User) UserLine {
	return UserLine{
		Username:        user.Username,
		Password:        user.Password,
		ActiveSessionId: user.ActiveSessionId,
	}
}

func lineToUser(id domain.UserId, userLine UserLine) domain.User {
	return domain.User{
		UserId:          id,
		Username:        userLine.Username,
		Password:        userLine.Password,
		ActiveSessionId: userLine.ActiveSessionId,
	}
}

func NewUser() *User {
	return &User{
		data: make(map[domain.UserId]UserLine),
	}
}

func (s *User) Post(user domain.User) error {
	if _, exists := s.data[user.UserId]; exists {
		return repository.AlreadyExists
	}
	s.data[user.UserId] = userToLine(user)
	return nil
}

func (s *User) Put(user domain.User) error {
	s.data[user.UserId] = userToLine(user)
	return nil
}

func (s *User) Get(id domain.UserId) (domain.User, error) {
	line, exists := s.data[id]
	if !exists {
		return domain.User{}, repository.NotFound
	}
	return lineToUser(id, line), nil
}

type Session struct {
	data map[domain.SessionId]domain.UserId
}

func NewSession() *Session {
	return &Session{
		make(map[domain.SessionId]domain.UserId),
	}
}

func (s *Session) Post(session domain.Session) error {
	if _, exists := s.data[session.SessionId]; exists {
		return repository.AlreadyExists
	}
	s.data[session.SessionId] = session.UserId
	return nil
}

func (s *Session) Put(session domain.Session) error {
	s.data[session.SessionId] = session.UserId
	return nil
}

func (s *Session) Get(id domain.SessionId) (domain.Session, error) {
	user, exists := s.data[id]
	if !exists {
		return domain.Session{}, repository.NotFound
	}
	return domain.Session{SessionId: id, UserId: user}, nil
}

func (s *Session) Delete(id domain.SessionId) error {
	if _, exists := s.data[id]; !exists {
		return repository.NotFound
	}
	delete(s.data, id)
	return nil
}
