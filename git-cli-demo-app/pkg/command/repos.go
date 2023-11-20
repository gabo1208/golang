package command

import (
	"fmt"

	"example.com/git-cli/pkg/packages"
	"github.com/spf13/cobra"
)

// NewVersionsCommand implements 'git-cli package-version' command
func NewVersionsCommand() *cobra.Command {
	var packageVersionsCmd = &cobra.Command{
		Use:   "package-version",
		Short: "Get pinned package version from git",
		Long: `Get pinned package version from git.
	If --targetRepoDir is set the CLI will
	- Check for internal deps or bundles versions
	If --targetRepoDir and --targetRepo are set the CLI will try to:
	- Clone the repo if it does not exist`,
		Run: func(cmd *cobra.Command, args []string) {
			if !silent {
				fmt.Println("Running git-cli package-version")
			}
			packages.GetPackageVersion(version, targetRepoDir, targetRepo, silent)
		},
	}
	// Set packageVersionsCmd options
	silentOption(packageVersionsCmd)
	versionOption(packageVersionsCmd)
	targetRepoDirOption(packageVersionsCmd)
	targetRepoOption(packageVersionsCmd)
	return packageVersionsCmd
}
