package gomodcp

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Prüft, ob im Projekt-Root irgendeine .go-Datei
// `package main` verwendet.
// → In diesem Fall darf KEIN Root-Package umgeschrieben werden.
func rootHasMainPackage(root string) (bool, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return false, err
	}

	re := regexp.MustCompile(`(?m)^package\s+main\b`)

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		src, err := os.ReadFile(filepath.Join(root, e.Name()))
		if err != nil {
			return false, err
		}

		if re.Match(src) {
			return true, nil
		}
	}

	return false, nil
}

// Leitet einen Go-konformen Root-Package-Namen
// aus dem Modulnamen ab.
// Bindestriche sind in Go-Packages nicht erlaubt.
func rootPackageFromModule(module string) string {
	parts := strings.Split(module, "/")
	last := parts[len(parts)-1]

	return strings.ReplaceAll(last, "-", "")
}

// Ersetzt `package xyz` in allen .go-Dateien
// DIREKT im Projekt-Root (keine Unterordner!)
func rewriteRootPackages(root string, newPkg string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(?m)^package\s+\w+`)

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		path := filepath.Join(root, e.Name())

		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		updated := re.ReplaceAll(
			src,
			[]byte("package "+newPkg),
		)

		if !bytes.Equal(src, updated) {
			if err := os.WriteFile(path, updated, 0644); err != nil {
				return err
			}
		}
	}

	return nil
}

// Entscheidet, ob das Root-Package umgeschrieben werden darf
// und tut es ggf.
func maybeRewriteRootPackage(projectRoot, module string) error {
	hasMain, err := rootHasMainPackage(projectRoot)
	if err != nil {
		return err
	}

	// Sonderfall: Root ist ein Programm
	if hasMain {
		return nil
	}

	rootPkg := rootPackageFromModule(module)
	return rewriteRootPackages(projectRoot, rootPkg)
}
