/*
Copyright © 2023 Gabriel Freites <gabohd@gmail.com>
*/
package main

import (
	"fmt"
	"os"

	"example.com/git-cli/pkg/root"
)

func main() {
	err := root.NewRootCommand().Execute()
	if err != nil {
		if err.Error() != "subcommand is required" {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
