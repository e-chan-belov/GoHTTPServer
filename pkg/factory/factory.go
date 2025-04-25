package factory

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"hw2/api/http"
	"hw2/repository"
	"hw2/repository/ram_storage"
	"hw2/usecases/service"
)

func NewServer() *chi.Mux {
	repo := repository.Task(ram_storage.NewTask())
	userRepo := repository.User(ram_storage.NewUser())
	sessionRepo := repository.Session(ram_storage.NewSession())

	taskService := service.NewTask(&repo)
	userManager := service.NewUserManager(&userRepo, &sessionRepo)
	taskHandlers := http.NewTaskHandler(taskService, userManager)
	userHandler := http.NewUserHandler(userManager)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	taskHandlers.WithTaskHandlers(r)
	userHandler.WithUserHandler(r)
	return r
}
