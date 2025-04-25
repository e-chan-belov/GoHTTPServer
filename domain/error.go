package domain

type ErrorMessage struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}
