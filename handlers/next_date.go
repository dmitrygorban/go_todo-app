package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dmitrygorban/go_todo-app/scheduler"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	nowTime, err := time.Parse("20060102", now)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error %s", err)
		return
	}
	next, err := scheduler.NextDate(nowTime, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(next))
}
