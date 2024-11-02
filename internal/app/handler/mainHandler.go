package handler

import "net/http"

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's my app created for example golang architecture"))
}