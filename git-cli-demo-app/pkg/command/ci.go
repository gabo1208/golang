package command

import (
	"fmt"

	"example.com/git-cli/pkg/ci"
	"github.com/spf13/cobra"
)

// NewCICommand implements 'git-cli package-version' command
func NewCICommand() *cobra.Command {
	var packageCICmd = &cobra.Command{
		Use:   "package-ci",
		Short: "Get package ci in git",
		Run: func(cmd *cobra.Command, args []string) {
			if !silent {
				fmt.Println("Running git-cli package-ci")
			}
			ci.GetPackageCIInfo(version, targetRepoDir, targetRepo, silent)
		},
	}
	// Set packageVersionsCmd options
	silentOption(packageCICmd)
	versionOption(packageCICmd)
	targetRepoDirOption(packageCICmd)
	targetRepoOption(packageCICmd)
	return packageCICmd
}
