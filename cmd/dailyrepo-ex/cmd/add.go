package cmd

import (
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	_ "github.com/taxio/go-cli-tutorial/statik"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "日報を生成",
	Long:  `今日の日付の日報を作成する`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		err = generateReport(fileName)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", time.Now().Format("2006-01-02")+".md", "report name")
}

func generateReport(filename string) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}
	// template読み込む
	tplFile, err := statikFs.Open("/report.md.tmpl")
	if err != nil {
		return err
	}
	defer tplFile.Close()
	btpl, err := ioutil.ReadAll(tplFile)
	if err != nil {
		return err
	}
	stpl := string(btpl)
	tpl := template.Must(template.New("report").Parse(stpl))

	// Todayを差し込んでファイルに書き込む
	rptFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	rptData := struct {
		Today string
	}{
		Today: time.Now().Format("2006-01-02"),
	}
	err = tpl.Execute(rptFile, rptData)
	if err != nil {
		return err
	}

	return nil
}
