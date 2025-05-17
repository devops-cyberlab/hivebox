// File: git-handler/tagging/tag.go
package tagging

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetCurrentGitTag returns the latest Git tag or commit hash
func GetCurrentGitTag() (string, error) {
	// Try to get the latest tag
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err == nil {
		return strings.TrimSpace(string(output)), nil
	}

	// Fallback: Get short commit hash
	cmd = exec.Command("git", "rev-parse", "--short", "HEAD")
	output, err = cmd.Output()
	if err == nil {
		return strings.TrimSpace(string(output)), nil
	}

	return "", fmt.Errorf("failed to get Git version")
}
