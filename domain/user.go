package domain

type UserId string
type UserLogin string
type UserPassword string
type SessionId string

type User struct {
	UserId          UserId       `json:"userId"`
	Username        UserLogin    `json:"username"`
	Password        UserPassword `json:"password"`
	ActiveSessionId SessionId    `json:"activeSessionId"`
}

type Session struct {
	SessionId SessionId `json:"session_id"`
	UserId    UserId    `json:"user_id"`
}
