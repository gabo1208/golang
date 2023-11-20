package ci

import (
	"fmt"
	"os"

	"example.com/git-cli/pkg/packages"
	"example.com/git-cli/pkg/repos"
	"example.com/git-cli/pkg/validations"
)

var Silent bool

func GetPackageCIInfo(version, targetRepoDir, targetRepo string, silent bool) {
	var err error
	var initialDir string
	Silent = silent
	repos.Silent = silent
	packages.Silent = silent

	if err = validations.ValidateVersion(version); err != nil {
		fmt.Println(err)
		return
	}
	if initialDir, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}

	// Get the target repo
	if err = repos.GetPackageRepo(targetRepoDir, targetRepo, "0", initialDir); err != nil {
		fmt.Println(err, validations.AddFriendlyErrorMessage(err))
		return
	}
	if !Silent {
		fmt.Println()
	}
	fmt.Printf("ci info for %s\n", targetRepo)
}
