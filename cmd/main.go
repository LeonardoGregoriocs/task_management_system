package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leonardogregoriocs/task_management_system/internal/config"
	"github.com/leonardogregoriocs/task_management_system/internal/database"
	"github.com/leonardogregoriocs/task_management_system/internal/endpoints"
	"github.com/leonardogregoriocs/task_management_system/internal/repository"
	"github.com/leonardogregoriocs/task_management_system/internal/user"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cfg := config.NewConfig()

	db, err := database.Connect(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// Todo - RunMigrations para DB.

	userService := user.UserService{
		Repository: &repository.UserRepository{Db: db},
	}

	handler := endpoints.Handler{
		UserService: userService,
	}

	// User
	r.Post("/user", handler.CreateUser)         // Create new user
	r.Get("/users", handler.GetAll)             // Get users
	r.Get("/user/{cpf}", handler.GetUserByCFP)  // Get user by cpf
	r.Put("/user/{cpf}", handler.UpdateUser)    // Update user
	r.Delete("/user/{cpf}", handler.DeleteUser) // Delete user

	http.ListenAndServe(":9001", r)
}
