package server

import (
	"log"
	"net/http"
)

// StaticServer se
type StaticServer struct{}

// Serve sds
func (staticServer *StaticServer) Serve() error {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	log.Println("Listening on :80...")
	return http.ListenAndServe(":80", nil)
}
