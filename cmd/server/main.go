package main

import (
	"net/http"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/server"
)

func main() {
	main := v1.NewTabloServiceServer(server.Tablo(), nil)

	mux := http.NewServeMux()
	mux.Handle(main.PathPrefix(), main)

	_ = http.ListenAndServe(":8080", mux)
}
