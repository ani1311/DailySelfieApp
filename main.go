package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

var (
	repoOwner = "ani1311"
	repoName  = "DailySelfieApp"
)

func main() {
	fs := http.FS(staticFiles)
	http.Handle("/", http.FileServer(fs))

	http.HandleFunc("/upload", uploadhander)
	http.HandleFunc("/ping", pingHandler)

	http.ListenAndServe(":8091", nil)
}
