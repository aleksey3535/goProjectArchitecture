package handler

import "net/http"

func (h *Handler) HandlerPartTwo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handler2 do something"))
}