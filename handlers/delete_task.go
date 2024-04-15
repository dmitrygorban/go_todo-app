package handlers

import (
	"net/http"
	"strconv"
)

func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
	queryId := r.FormValue("id")

	idInt, err := strconv.ParseInt(queryId, 10, 0)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Store.Delete(int(idInt))
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = emptyOkResponse(w)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
