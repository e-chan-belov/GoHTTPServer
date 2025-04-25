package types

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"hw2/api"
	"hw2/domain"
	"net/http"
)

type PostTaskHandlerRequest struct {
	domain.TaskName `json:"task"`
}

func CreatePostTaskHandlerRequest(r *http.Request) (*PostTaskHandlerRequest, error) {
	var req PostTaskHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

type PostTaskHandlerAnswer struct {
	domain.TaskId `json:"task_id"`
}

func AnswerPostTaskHandler(w http.ResponseWriter, id domain.TaskId) {
	ans := PostTaskHandlerAnswer{id}
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		ProcessError(w, 500, "types", err)
		return
	}
}

type GetStatusHandlerRequest struct {
	domain.TaskId
}

func CreateGetStatusHandlerRequest(r *http.Request) (*GetStatusHandlerRequest, error) {
	var req = GetStatusHandlerRequest{domain.TaskId(chi.URLParam(r, "id"))}
	if req.TaskId == "" {
		return nil, api.NoId
	}
	return &req, nil
}

type GetStatusHandlerAnswer struct {
	domain.TaskStatus `json:"task_status"`
}

func AnswerGetStatusHandler(w http.ResponseWriter, status domain.TaskStatus) {
	ans := GetStatusHandlerAnswer{status}
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		ProcessError(w, 500, "types", err)
		return
	}
}

type GetResultHandlerRequest struct {
	domain.TaskId
}

func CreateGetResultHandlerRequest(r *http.Request) (*GetResultHandlerRequest, error) {
	var req = GetResultHandlerRequest{domain.TaskId(chi.URLParam(r, "id"))}
	if req.TaskId == "" {
		return nil, api.NoId
	}
	return &req, nil
}

type GetResultHandlerAnswer struct {
	domain.TaskResult `json:"task_result"`
}

func AnswerGetResultHandler(w http.ResponseWriter, result domain.TaskResult) {
	ans := GetResultHandlerAnswer{result}
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		ProcessError(w, 500, "types", err)
		return
	}
}
