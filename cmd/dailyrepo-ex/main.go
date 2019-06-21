package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/taxio/go-cli-tutorial/cmd/dailyrepo-ex/cmd"
)

func main() {
	fs := afero.NewOsFs()
	out := os.Stdout

	if err := cmd.NewRootCmd(out, fs).Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
