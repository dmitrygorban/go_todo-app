package handlers

import (
	"net/http"
)

func (h *TaskHandler) HandleTaskRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.add(w, r)
		return
	default:
		handleError(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		return
	}
}
