package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_printVersion(t *testing.T) {
	out := new(bytes.Buffer)
	err := printVersion(out)
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("dailyrepo-ex %s\n", version)
	if want != out.String() {
		t.Fatalf("printVersion() output %s, but want %s\n", out.String(), want)
	}
}
