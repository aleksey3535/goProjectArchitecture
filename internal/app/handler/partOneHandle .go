package handler

import "net/http"

func (h *Handler) HandlerPartOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handler1 do something"))
}