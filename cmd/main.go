package main

import (
	"context"
	"fmt"
	"os"

	"github.com/axelrhd/gomodcp"
)

func main() {
	cmd := gomodcp.NewCLI()

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		// Klare Fehlermeldung für den User
		fmt.Fprintln(os.Stderr, "error:", err)

		// Optional: kurzer Usage-Hinweis
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Run 'gomodcp --help' for more information.")

		os.Exit(1)
	}
}
