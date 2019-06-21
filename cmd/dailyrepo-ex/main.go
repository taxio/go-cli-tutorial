package main

import (
	"log"

	"github.com/spf13/afero"
	"github.com/taxio/go-cli-tutorial/cmd/dailyrepo-ex/cmd"
)

func main() {
	fs := afero.NewOsFs()
	if err := cmd.NewRootCmd(fs).Execute(); err != nil {
		log.Fatal(err)
	}
}
