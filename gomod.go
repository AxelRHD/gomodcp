package gomodcp

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func readModuleName(goModPath string) (string, error) {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", errors.New("module declaration not found in go.mod")
}

func rewriteGoMod(dstDir, newModule string) error {
	path := filepath.Join(dstDir, "go.mod")

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "module ") {
			lines[i] = "module " + newModule
			break
		}
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644)
}
