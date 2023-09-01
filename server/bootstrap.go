package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

var router = mux.NewRouter()

func AddHandler(path string, handler http.Handler) {
	router.Handle(path, handler)
}

func Start(port string) {
	log.Println("Listening on: " + port)
	router.Handle("/live.js", live.Javascript{})
	http.Handle("/", router)
	http.ListenAndServe(port, nil)
}
