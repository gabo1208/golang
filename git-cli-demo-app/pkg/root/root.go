package root

import (
	"example.com/git-cli/pkg/command"
	"github.com/spf13/cobra"
)

// NewRootCommand represents the plugin's entrypoint
func NewRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "git-cli",
		Short: "git CLI tool.",
		Long:  `Simplify and automate Git tasks with go.`,
	}

	rootCmd.AddCommand(command.NewVersionsCommand())
	rootCmd.AddCommand(command.NewCICommand())
	return rootCmd
}
