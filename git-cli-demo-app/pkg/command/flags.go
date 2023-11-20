package command

import (
	"github.com/spf13/cobra"
)

var (
	silent bool

	version       string
	targetRepo    string
	targetRepoDir string
)

func silentOption(targetCmd *cobra.Command) {
	targetCmd.Flags().BoolVarP(
		&silent,
		"silent",
		"s",
		false,
		"(optional) Print out just required output")
}

func versionOption(targetCmd *cobra.Command) {
	targetCmd.Flags().StringVar(
		&version,
		"version",
		"1.0",
		"The version",
	)
}

func targetRepoDirOption(targetCmd *cobra.Command) {
	targetCmd.Flags().StringVar(
		&targetRepoDir,
		"target-repo-dir",
		"test",
		"(optional) The path to search/clone the target repo")
}

func targetRepoOption(targetCmd *cobra.Command) {
	targetCmd.Flags().StringVar(
		&targetRepo,
		"target-repo",
		"git@github.com:your/repo.git",
		"(optional) The target repo ssh url")
}
