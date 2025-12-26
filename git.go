package gomodcp

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func gitTrackedFiles(dir string) ([]string, error) {
	cmd := exec.Command("git", "-C", dir, "ls-files")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git ls-files failed (is this a git repo?): %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return nil, errors.New("no git-tracked files found")
	}

	return lines, nil
}
