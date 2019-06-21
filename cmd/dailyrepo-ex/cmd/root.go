package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var verbose = false
var version = "v0.0.1"

func NewRootCmd(out io.Writer, fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dailyrepo-ex",
		Short: "日報作成ツール",
		Long:  `テンプレートから日報を作成します．`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// version
			v, err := cmd.Flags().GetBool("version")
			if err != nil {
				return err
			}
			if v {
				return printVersion(out)
			}

			return nil
		},
	}

	cmd.Flags().BoolP("version", "v", false, "print version")
	cmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "print debug log")

	cmd.AddCommand(NewAddCmd(fs))

	return cmd
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "dailyrepo-ex %s\n", version)
	return err
}
