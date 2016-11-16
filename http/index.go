package http

import (
	"github.com/kerwindena/overtime"

	"net/http"
)

func index(srv *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		size := 1000
		dates := make([]overtime.Date, size, size)
		for i := 0; i < size; i++ {
			dates[i] = srv.dateStorage.Date(17500 - size/2 + i)
		}
		srv.templates.ExecuteTemplate(w, "default.html", dates)
	}
}
