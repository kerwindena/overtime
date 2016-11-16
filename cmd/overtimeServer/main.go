package main

import (
	"github.com/kerwindena/overtime"
	"github.com/kerwindena/overtime/http"
	"github.com/kerwindena/overtime/time"
)

func main() {
	var ds overtime.DateStorage
	ds = time.NewDateStorage()

	server := http.NewServer(ds)

	server.ListenAndServe()
}
