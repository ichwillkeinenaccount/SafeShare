package api

import "net/http"

type TextHandler struct {
}

func (h *TextHandler) getAll(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		return
	}
}

func (h *TextHandler) postText(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		return
	}
}
