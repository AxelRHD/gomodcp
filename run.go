package gomodcp

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run(srcDir, destMod, dstDir string, useGit bool) error {
	srcAbs, err := filepath.Abs(srcDir)
	if err != nil {
		return err
	}

	dstAbs, err := filepath.Abs(dstDir)
	if err != nil {
		return err
	}

	if _, err := os.Stat(dstAbs); err == nil {
		return fmt.Errorf("destination already exists: %s", dstAbs)
	}

	oldModule, err := readModuleName(filepath.Join(srcAbs, "go.mod"))
	if err != nil {
		return err
	}

	if useGit {
		if err := copyTreeGit(srcAbs, dstAbs); err != nil {
			return err
		}
	} else {
		if err := copyTreeFS(srcAbs, dstAbs); err != nil {
			return err
		}
	}

	if err := rewriteGoMod(dstAbs, destMod); err != nil {
		return err
	}

	return rewriteImports(dstAbs, oldModule, destMod)
}
