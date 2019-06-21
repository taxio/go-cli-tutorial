package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/rakyll/statik/fs"
	_ "github.com/taxio/go-cli-tutorial/statik"
)

type RootOption struct {
	flgVersion bool
	flgVerbose bool
}

type AddOption struct {
	fileName string
}

func main() {
	rootOpt := RootOption{}
	rootCmd := flag.NewFlagSet("dailyrepo root", flag.ContinueOnError)
	rootCmd.BoolVar(&rootOpt.flgVersion, "v", false, "print version(short)")
	rootCmd.BoolVar(&rootOpt.flgVersion, "version", false, "print version")
	rootCmd.BoolVar(&rootOpt.flgVerbose, "verbose", false, "print info log for developer")
	args := os.Args[1:]
	err := rootCmd.Parse(args)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
	err = handleRootCmd(&rootOpt)
	if err != nil {
		log.Fatal(err)
	}

	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)
	addOpt := AddOption{}
	defaultFileName := time.Now().Format("2006-01-02") + ".md"
	addCmd.StringVar(&addOpt.fileName, "name", defaultFileName, "specify generating filename")

	args = rootCmd.Args()
	if len(args) > 0 {
		switch args[0] {
		case "add":
			err := addCmd.Parse(args[1:])
			if err != nil {
				if err == flag.ErrHelp {
					os.Exit(0)
				}
				log.Fatal(err)
			}
			args = addCmd.Args()
			err = handleAddCmd(&addOpt)
			if err != nil {
				log.Fatal(err)
			}

		default:
			rootCmd.Usage()
			os.Exit(2)
		}
	}
}

func handleRootCmd(opt *RootOption) error {
	if opt.flgVersion {
		fmt.Println("v0.0.1")
	}
	return nil
}

func handleAddCmd(opt *AddOption) error {
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
	rptFile, err := os.Create(opt.fileName)
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
