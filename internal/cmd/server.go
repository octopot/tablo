package cmd

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/spf13/cobra"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	server "go.octolab.org/ecosystem/tablo/internal/server/v1"
)

// New returns the new server command.
func NewServer() *cobra.Command {
	command := cobra.Command{
		Use:   "tablo",
		Short: "the one point of view to all your task boards",
		Long:  "The one point of view to all your task boards.",

		SilenceErrors: false,
		SilenceUsage:  true,

		RunE: func(cmd *cobra.Command, args []string) error {
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

			return http.ListenAndServe(":8080", mux)
		},
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}
