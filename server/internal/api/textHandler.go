package api

import "net/http"

type TextHandler struct {
}

// getAll returns all texts
// @Summary      Show all texts
// @Description  get all texts
// @Tags         text
// @Produce      json
// @Success      200  string string
// @Failure      400  string string
// @Failure      404  string string
// @Failure      500  string string
// @Router       /api/v1/text/ [get]
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
