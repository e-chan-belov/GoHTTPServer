package types

import (
	"encoding/json"
	"hw2/api"
	"hw2/domain"
	"net/http"
)

func GetAuthToken(r *http.Request) (domain.SessionId, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return domain.SessionId(token), api.NoAuth
	}
	return domain.SessionId(token), nil
}

type UserBodyRequest struct {
	domain.UserLogin    `json:"login"`
	domain.UserPassword `json:"password"`
}

func CreateUserBodyRequest(r *http.Request) (UserBodyRequest, error) {
	var req UserBodyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return UserBodyRequest{}, err
	}
	return req, nil
}

type LoginRequestAnswer struct {
	domain.SessionId `json:"sessionId"`
}

func CreateAnswerLoginRequest(w http.ResponseWriter, token domain.SessionId) {
	ans := LoginRequestAnswer{token}
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		ProcessError(w, 500, "types", err)
		return
	}
}
