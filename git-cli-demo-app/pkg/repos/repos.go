package repos

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"example.com/git-cli/pkg/validations"
	"gopkg.in/yaml.v3"
)

var Silent bool

type PackageMetadataYaml struct {
	Name string `yaml:"name"`
}

type PackageYaml struct { // partial parsing for the package.yml file
	Metadata PackageMetadataYaml `yaml:"metadata"`
}

type PackageVersionYaml struct {
	Version map[string]string `yaml:"version"`
}

func GoBackToScriptRootDir(scriptDir string) { // always return to the scripts dir
	if err := os.Chdir(scriptDir); err != nil {
		fmt.Println(err)
	}
}

func runCommandWithOutput(command *exec.Cmd) error {
	if !Silent {
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}
	return command.Run()
}

func ExecGitCommandInDir(repoDir, initialDir string, command *exec.Cmd) error {
	defer GoBackToScriptRootDir(initialDir)
	if err := os.Chdir(repoDir); err != nil { // move to the repo dir
		return err
	}
	if err := runCommandWithOutput(command); err != nil {
		return err
	}
	return nil
}

func GetGitBranchesList(repoDir, initialDir, version string) ([]string, error) {
	defer GoBackToScriptRootDir(initialDir)
	if err := os.Chdir(repoDir); err != nil { // move to the repo dir
		return nil, err
	}
	stdOut, err := exec.Command("git", "branch", "-r", "--list", fmt.Sprintf("origin/%s.[0-9]*", version)).Output()
	if err != nil {
		return nil, err
	}

	branches := []string{} // clean the branches output array
	for _, branch := range strings.Split(string(stdOut), "\n") {
		trimmedBranch := strings.TrimSpace(branch)
		if strings.HasPrefix(trimmedBranch, "origin/") {
			branches = append(branches, trimmedBranch)
		}
	}
	return branches, nil
}

func GetGitCurrentBranch(repoDir, initialDir string) (string, error) {
	defer GoBackToScriptRootDir(initialDir)
	if err := os.Chdir(repoDir); err != nil { // move to the repo dir
		return "", err
	}
	stdOut, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(stdOut), "\n"), nil
}

func ParseYamlFile(filePath string, parseBuff interface{}) error {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(buf), parseBuff)
	if err != nil {
		return err
	}
	return nil
}

func GetPackageRepo(packageDir, packageRepo, depth, initialDir string) error {
	exists, err := validations.ValidateDirExists(packageDir)
	if err != nil {
		return err
	}

	if exists { // if path exists try to open repo
		if !Silent {
			fmt.Printf("%s dir exists, using local repo\n", packageDir)
		}
		if err := ExecGitCommandInDir(packageDir, initialDir, exec.Command("git", "status")); err != nil {
			return err
		}
	} else { // if path does not exists try to clone the repo
		if packageRepo == "" {
			return errors.New("no repo URL set, error trying to clone the repo")
		}

		cloneCommand := exec.Command("git", "clone", packageRepo, packageDir)
		if depth != "0" {
			cloneCommand = exec.Command("git", "clone", "--depth", depth, packageRepo, packageDir)
		}
		// if path does not exists try to clone repo
		err := runCommandWithOutput(cloneCommand)
		if err != nil {
			return err
		}
	}

	if err := ExecGitCommandInDir(packageDir, initialDir, exec.Command("git", "fetch", "--tags", "--all")); err != nil { // fetch repo
		return err
	}
	return nil
}
