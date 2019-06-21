package cmd

import (
	"io/ioutil"
	"text/template"
	"time"

	"github.com/rakyll/statik/fs"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	_ "github.com/taxio/go-cli-tutorial/statik"
)

func NewAddCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "日報を生成",
		Long:  `今日の日付の日報を作成する`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fileName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			err = generateReport(fs, fileName)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringP("name", "n", time.Now().Format("2006-01-02")+".md", "report name")

	return cmd
}

func generateReport(afs afero.Fs, filename string) error {
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

	rptFile, err := afs.Create(filename)
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
