package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	_ "github.com/taxio/go-cli-tutorial/statik"
)

func Test_generateReport(t *testing.T) {
	fs := afero.NewMemMapFs()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	filename := filepath.Join(wd, "test.md")

	af := afero.Afero{Fs: fs}
	ext, err := af.Exists(filename)
	if err != nil {
		t.Fatal(err)
	}
	if ext {
		t.Fatal(filename, " already exists.")
	}

	err = generateReport(fs, filename)
	if err != nil {
		t.Fatal(err)
	}
	ext, err = af.Exists(filename)
	if err != nil {
		t.Fatal(err)
	}
	if !ext {
		t.Fatal(filename, " not exists.")
	}
}
