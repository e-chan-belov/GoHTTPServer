package http

import (
	"github.com/go-chi/chi"
	"hw2/api/http/types"
	"hw2/usecases"
	"net/http"
)

// Task represents an HTTP handler for managing tasks.
type Task struct {
	service     usecases.Task
	userManager usecases.UserManager
}

// NewTaskHandler creates a new instance of Task.
func NewTaskHandler(service usecases.Task, manager usecases.UserManager) *Task {
	return &Task{service: service,
		userManager: manager,
	}
}

// PostTaskHandler
// @Summary Create a task
// @Description Create a new task with the specified request
// @Tags object
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authorization token in format: Bearer {auth_token}"
// @Param request body types.PostTaskHandlerRequest true "Task request"
// @Success 200 {object} types.PostTaskHandlerAnswer
// @Failure 400 {object} domain.ErrorMessage
// @Router /task [post]
func (s *Task) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	tok, err := types.GetAuthToken(r)
	if err != nil {
		types.ProcessError(w, 401, "PostTaskHandler! ", err)
		return
	}
	if err := s.userManager.Auth(tok); err != nil {
		types.ProcessError(w, 400, "PostTaskHandler! ", err)
		return
	}

	req, err := types.CreatePostTaskHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, 400, "PostTaskHandler! ", err)
		return
	}
	ans, err := s.service.PostTask(req.TaskName)
	if err != nil {
		types.ProcessError(w, 400, "PostTaskHandler! ", err)
		return
	}
	types.AnswerPostTaskHandler(w, ans)
}

// GetStatusHandler
// @Summary Get a task's status
// @Description Get a task's status by its id
// @Tags object
// @Accept plain
// @Produce json
// @Param Authorization header string true "Authorization token in format: Bearer {auth_token}"
// @Param id path string true "ID объекта"
// @Success 200 {object} types.GetStatusHandlerAnswer
// @Failure 400 {object} domain.ErrorMessage
// @Failure 404 {object} domain.ErrorMessage
// @Router /status/{id} [get]
func (s *Task) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	tok, err := types.GetAuthToken(r)
	if err != nil {
		types.ProcessError(w, 401, "PostTaskHandler! ", err)
		return
	}
	if err := s.userManager.Auth(tok); err != nil {
		types.ProcessError(w, 400, "PostTaskHandler! ", err)
		return
	}

	req, err := types.CreateGetStatusHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, 400, "GetStatusHandler error! ", err)
		return
	}
	ans, err := s.service.GetStatus(req.TaskId)
	if err != nil {
		types.ProcessError(w, 404, "GetStatusHandler error! ", err)
		return
	}
	types.AnswerGetStatusHandler(w, ans)
}

// GetResultHandler
// @Summary Get a task's result
// @Description Get a task's result by its id
// @Tags object
// @Accept plain
// @Produce json
// @Param Authorization header string true "Authorization token in format: Bearer {auth_token}"
// @Param id path string true "ID объекта"
// @Success 200 {object} types.GetResultHandlerAnswer
// @Failure 400 {object} domain.ErrorMessage
// @Failure 404 {object} domain.ErrorMessage
// @Router /result/{id} [get]
func (s *Task) GetResultHandler(w http.ResponseWriter, r *http.Request) {
	tok, err := types.GetAuthToken(r)
	if err != nil {
		types.ProcessError(w, 401, "PostTaskHandler! ", err)
		return
	}
	if err := s.userManager.Auth(tok); err != nil {
		types.ProcessError(w, 400, "PostTaskHandler! ", err)
		return
	}

	req, err := types.CreateGetResultHandlerRequest(r)
	if err != nil {
		types.ProcessError(w, 400, "GetResultHandler error! ", err)
		return
	}
	ans, err := s.service.GetResult(req.TaskId)
	if err != nil {
		types.ProcessError(w, 404, "GetResultHandler error! ", err)
		return
	}
	types.AnswerGetResultHandler(w, ans)
}

// WithTaskHandlers registers object-related HTTP handlers.
func (s *Task) WithTaskHandlers(r chi.Router) {
	r.Route("/task", func(r chi.Router) {
		r.Post("/", s.PostTaskHandler)
	})
	r.Route("/status/{id}", func(r chi.Router) {
		r.Get("/", s.GetStatusHandler)
	})
	r.Route("/result/{id}", func(r chi.Router) {
		r.Get("/", s.GetResultHandler)
	})
}
