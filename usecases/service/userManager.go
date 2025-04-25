package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"hw2/domain"
	"hw2/repository"
)

type UserManager struct {
	users    *repository.User
	sessions *repository.Session
}

func NewUserManager(users *repository.User, sessions *repository.Session) *UserManager {
	return &UserManager{
		users:    users,
		sessions: sessions,
	}
}

func (m *UserManager) Register(username domain.UserLogin, pass domain.UserPassword) error {
	hash := sha256.Sum256([]byte(username))
	id := hex.EncodeToString(hash[:])
	user := domain.User{UserId: domain.UserId(id), Username: username, Password: pass, ActiveSessionId: ""}
	if err := (*m.users).Post(user); err != nil {
		return err
	}
	return nil
}

func (m *UserManager) Login(username domain.UserLogin, pass domain.UserPassword) (domain.SessionId, error) {
	hash := sha256.Sum256([]byte(username))
	id := hex.EncodeToString(hash[:])
	user, err := (*m.users).Get(domain.UserId(id))
	if err != nil {
		return "", err
	}

	if user.ActiveSessionId != "" {
		if err := (*m.sessions).Delete(user.ActiveSessionId); err != nil {
			return "", err
		}
	}

	user.ActiveSessionId = domain.SessionId(uuid.New().String())
	if err := (*m.users).Put(user); err != nil {
		return "", err
	}
	if err := (*m.sessions).Post(domain.Session{SessionId: user.ActiveSessionId, UserId: user.UserId}); err != nil {
		return "", err
	}
	return user.ActiveSessionId, nil
}

func (m *UserManager) Auth(id domain.SessionId) error {
	session, err := (*m.sessions).Get(id)
	if err != nil {
		return err
	}
	user, err := (*m.users).Get(session.UserId)
	if err != nil {
		return err
	}
	if user.ActiveSessionId != id {
		return repository.WrongToken
	}
	return nil
}
