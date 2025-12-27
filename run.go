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

	// 1) go.mod umschreiben (neuer Modulname)
	if err := rewriteGoMod(dstAbs, destMod); err != nil {
		return err
	}

	// 2) 🆕 Root-Package ggf. umschreiben
	//    - nur Root-Ebene
	//    - NIE, wenn package main im Root existiert
	if err := maybeRewriteRootPackage(dstAbs, destMod); err != nil {
		return err
	}

	// 3) Imports anpassen
	return rewriteImports(dstAbs, oldModule, destMod)
}
