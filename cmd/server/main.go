package main

import (
	"log"
	"net/http"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/server"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

func main() {
	log.Printf("{commit: %q, date: %q, version: %q, port: 8080}\n", commit, date, version)

	main := v1.NewTabloServiceServer(server.Tablo(), nil)

	mux := http.NewServeMux()
	mux.Handle(main.PathPrefix(), main)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
