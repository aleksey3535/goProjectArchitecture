package app

import (
	"app/internal/app/handler"
	"app/internal/app/repository"
	"app/internal/app/service"
	"net/http"
	"time"
)

type App struct {
	h *handler.Handler
	s *service.Service
	r *repository.Repository
}

func New() (*App, error) {
	db, err := repository.NewPostgres(repository.Config{
		Host: "host",
		Port: "port",
		Username: "username",
		DBName: "dbname",
		Password: "password",
		SSLMode: "sslmode",
	})
	if err != nil {
		return nil, err
	}
	repos := repository.New(db)
	services := service.New(repos)
	handlers := handler.New(services)
	return &App{
		h: handlers,
		s: services,
		r: repos,
	}, nil
}

func(a *App) Run(port string) error {
	server := http.Server{
		Addr: ":" + port,
		Handler: a.h.InitRoutes(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server.ListenAndServe()

}