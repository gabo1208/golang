package validations

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ValidateVersion(version string) error {
	if version == "" {
		return errors.New("version must be set.")
	}
	return nil
}

func ValidateDirExists(dir string) (bool, error) {
	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if !stat.IsDir() {
		return true, errors.New("path is not a directory")
	}

	return true, nil
}

func AddFriendlyErrorMessage(err error) string {
	if strings.HasPrefix(err.Error(), "dial tcp:") {
		fmt.Println("Try checking your VPN connection.")
	}

	if strings.HasPrefix(err.Error(), "ssh:") {
		fmt.Println("Try checking your ssh keys configuration for git.")
	}

	return ""
}
