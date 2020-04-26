package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/cors"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	server "go.octolab.org/ecosystem/tablo/internal/server/v1"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

//nolint:gochecknoinits
func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
	}
}

func main() {
	log.Printf("{commit: %q, date: %q, version: %q, port: 8080}\n", commit, date, version)

	handler := v1.NewTabloServiceServer(server.Tablo(), nil)
	middleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	mux := http.NewServeMux()
	mux.Handle(handler.PathPrefix(), middleware.Handler(handler))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
