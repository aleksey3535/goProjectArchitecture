package handler

import (
	"app/internal/app/service"
	"net/http"
)

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s: s}
}

func(h *Handler) InitRoutes() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/",h.MainHandler )
	mux.HandleFunc("/part1/", h.HandlerPartOne)
	mux.HandleFunc("part2/", h.HandlerPartTwo)
	return mux
}
