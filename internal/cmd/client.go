package cmd

import "github.com/spf13/cobra"

// New returns the new client command.
func NewClient() *cobra.Command {
	command := cobra.Command{
		Use:   "tabloctl",
		Short: "the one point of view to all your task boards",
		Long:  "The one point of view to all your task boards.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}
