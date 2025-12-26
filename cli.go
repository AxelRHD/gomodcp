package gomodcp

import (
	"context"
	"fmt"
	"path/filepath"

	cli "github.com/urfave/cli/v3"
)

// Build-time variables (set via -ldflags)
var (
	appVersion = "dev"
	gitVersion = "unknown"
)

func NewCLI() *cli.Command {
	return &cli.Command{
		Name:      "gomodcp",
		Usage:     "copy & rename a Go project locally, updating module path and imports",
		ArgsUsage: "<src-dir> <dest-mod>",
		Description: `
<src-dir>
    Path to the source Go project directory.
    This directory must contain a go.mod file.

<dest-mod>
    Target Go module path.
    Example: github.com/axelrhd/myapp
`,
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "version",
				Usage: "print version information and exit",
			},
			&cli.StringFlag{
				Name:  "dst",
				Usage: "destination directory (optional)",
			},
			&cli.BoolFlag{
				Name:  "git",
				Usage: "use only git-tracked files (respects .gitignore)",
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			// --version short-circuit
			if c.Bool("version") {
				fmt.Printf("gomodcp v%s (%s)\n", appVersion, gitVersion)
				return nil
			}

			argCount := c.Args().Len()

			switch argCount {
			case 0:
				return fmt.Errorf("missing <src-dir> and <dest-mod>")
			case 1:
				return fmt.Errorf("missing <dest-mod>")
			case 2:
				// ok
			default:
				return fmt.Errorf("too many arguments (expected <src-dir> <dest-mod>)")
			}

			srcDir := c.Args().Get(0)
			destMod := c.Args().Get(1)

			dst := c.String("dst")
			if dst == "" {
				dst = filepath.Base(destMod)
			}

			return Run(srcDir, destMod, dst, c.Bool("git"))
		},
	}
}
