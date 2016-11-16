package http

import (
	"github.com/kerwindena/overtime"

	"html/template"
	"net"
	"net/http"
)

type Server struct {
	dateStorage overtime.DateStorage
	server      http.Server
	templates   *template.Template
}

func NewServer(dateStorage overtime.DateStorage) *Server {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	srv := &Server{
		server:      server,
		dateStorage: dateStorage,
		templates:   template.Must(template.ParseGlob("templates/*.html")),
	}

	mux.HandleFunc("/", index(srv))

	return srv
}

func (srv *Server) ListenAndServe() error {
	return srv.server.ListenAndServe()
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
	return srv.server.ListenAndServeTLS(certFile, keyFile)
}

func (srv *Server) Serve(l net.Listener) error {
	return srv.server.Serve(l)
}

func (srv *Server) SetKeepAlivesEnabled(v bool) {
	srv.server.SetKeepAlivesEnabled(v)
}
