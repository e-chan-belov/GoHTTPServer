package http

import (
	"github.com/go-chi/chi"
	"hw2/api/http/types"
	"hw2/usecases"
	"net/http"
)

// User represents an HTTP handler for user registration and creating sessions.
type User struct {
	userManager usecases.UserManager
}

// NewUserHandler creates a new instance of User.
func NewUserHandler(userManager usecases.UserManager) *User {
	return &User{
		userManager: userManager,
	}
}

// RegisterHandler
// @Summary Create new user and login
// @Description Creates new user and session for that user and gets a token for that user
// @Tags object
// @Accept json
// @Produce json
// @Param request body types.UserBodyRequest true "User request"
// @Success 200 {object} types.LoginRequestAnswer
// @Failure 400 {object} domain.ErrorMessage
// @Failure 401 {object} domain.ErrorMessage
// @Router /register [post]
func (s *User) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := types.CreateUserBodyRequest(r)
	if err != nil {
		types.ProcessError(w, 400, "RegisterHandler!", err)
		return
	}
	if err := s.userManager.Register(body.UserLogin, body.UserPassword); err != nil {
		types.ProcessError(w, 401, "RegisterHandler!", err)
		return
	}
	token, err := s.userManager.Login(body.UserLogin, body.UserPassword)
	if err != nil {
		types.ProcessError(w, 401, "RegisterHandler!", err)
		return
	}
	types.CreateAnswerLoginRequest(w, token)
}

// LoginHandler
// @Summary Get a login token
// @Description Creates new session for given user and gets a token for that user
// @Tags object
// @Accept json
// @Produce json
// @Param request body types.UserBodyRequest true "User request"
// @Success 200 {object} types.LoginRequestAnswer
// @Failure 400 {object} domain.ErrorMessage
// @Failure 401 {object} domain.ErrorMessage
// @Router /login [post]
func (s *User) LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := types.CreateUserBodyRequest(r)
	if err != nil {
		types.ProcessError(w, 400, "RegisterHandler!", err)
		return
	}
	token, err := s.userManager.Login(body.UserLogin, body.UserPassword)
	if err != nil {
		types.ProcessError(w, 401, "RegisterHandler!", err)
		return
	}
	types.CreateAnswerLoginRequest(w, token)
}

// WithUserHandler registers object-related HTTP handlers.
func (s *User) WithUserHandler(r chi.Router) {
	r.Route("/register", func(r chi.Router) {
		r.Post("/", s.RegisterHandler)
	})
	r.Route("/login", func(r chi.Router) {
		r.Post("/", s.LoginHandler)
	})
}
