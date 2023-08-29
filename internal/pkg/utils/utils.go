package utils

import (
	"os"
	"strings"
)

// Check currently only checks for the filename starting with "~",
// and tries to replace this with the current user's home directory.
func Check(filename string) (string, error) {
	if strings.HasPrefix(filename, "~/") {
		uhd, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return strings.ReplaceAll(filename, "~", uhd), nil
	}

	return filename, nil
}
