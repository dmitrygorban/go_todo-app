package handlers

import (
	"net/http"
)

func (h *TaskHandler) HandleTaskRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.add(w, r)
		return
	case http.MethodGet:
		h.getById(w, r)
		return
	case http.MethodPut:
		h.edit(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		handleError(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		return
	}
}
