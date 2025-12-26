package gomodcp

import (
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func rewriteImports(root, oldModule, newModule string) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return err
		}

		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		changed := false
		for _, imp := range file.Imports {
			val := strings.Trim(imp.Path.Value, `"`)
			if strings.HasPrefix(val, oldModule) {
				imp.Path.Value = `"` + strings.Replace(val, oldModule, newModule, 1) + `"`
				changed = true
			}
		}

		if !changed {
			return nil
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()

		cfg := &printer.Config{
			Mode:     printer.UseSpaces | printer.TabIndent,
			Tabwidth: 8,
		}
		return cfg.Fprint(f, fset, file)
	})
}
