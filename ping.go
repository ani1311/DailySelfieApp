package main

import (
	"net/http"

	"fortio.org/log"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Received ping request")
	w.Write([]byte("pong"))
}
