package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose = false
var version = "v0.0.1"

var rootCmd = &cobra.Command{
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
			fmt.Printf("dailyrepo-ex %s\n", version)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(func() {
	//	if !verbose{
	//		return
	//	}
	//	// ここでログの初期設定をしたりする
	//})

	rootCmd.Flags().BoolP("version", "v", false, "print version")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "print debug log")
}
