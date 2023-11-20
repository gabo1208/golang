package packages

import (
	"fmt"
	"os"
	"strings"

	"example.com/git-cli/pkg/repos"
	"example.com/git-cli/pkg/validations"
)

var Silent bool

func GetPackageVersion(version, targetRepoDir, targetRepo string, silent bool) {
	var err error
	var initialDir string
	checkTargetRepoDepsVersion := targetRepoDir != ""
	Silent = silent
	repos.Silent = silent

	if err = validations.ValidateVersion(version); err != nil {
		fmt.Println(err)
		return
	}
	if initialDir, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}

	// Get target package repo if --target-package-dir is set
	if checkTargetRepoDepsVersion {
		if err = repos.GetPackageRepo(targetRepoDir, targetRepo, "0", initialDir); err != nil {
			fmt.Println(err, validations.AddFriendlyErrorMessage(err))
			return
		}
	}

	var versionBranches []string
	if strings.Count(version, ".") == 2 {
		versionBranches = append(versionBranches, version)
	} else {
		versionBranches = append(versionBranches, version)
		branches, err := repos.GetGitBranchesList(targetRepoDir, initialDir, version)
		if err != nil {
			fmt.Println(err)
			return
		}

		latestBranch := strings.Split(branches[len(branches)-1], "/")
		if len(latestBranch) == 0 {
			fmt.Printf("No branches found for version: %s\n", version)
			return
		}
		versionBranches = append(versionBranches, latestBranch[len(latestBranch)-1])
	}
}
