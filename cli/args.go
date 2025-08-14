// Package cli
package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var ErrWithUsage = fmt.Errorf("usage printed")

type Conf struct {
	SourceJSONFilePath string
	SelectedOpt        string
	DefaultExec        string
}

func ParseArgs() (*Conf, error) {
	cfg := &Conf{}

	flag.Usage = func() {
		headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
		fmt.Fprintf(os.Stderr, "\n%s\n\n", headerFmt("Usage of %s:", os.Args[0]))

		tbl := table.New("Flag", "Type", "Description", "Default")
		tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
		tbl.AddRow("-h, --help", "boolean", "Print this message.", "")
		tbl.AddRow("-m, --menu-config", "string", "Path to the JSON menu configuration file. (Mandatory)", "")
		tbl.AddRow("-r, --picker-result", "string", "String coming back from the picker used to find relevant exec command in the config", "")
		tbl.AddRow("-e, --default-exec", "string", "Command to run on every result.", "")
		tbl.Print()
	}

	flag.StringVar(&cfg.SourceJSONFilePath, "m", "", "Path to the JSON menu configuration file.")
	flag.StringVar(&cfg.SourceJSONFilePath, "menu-config", "", "Path to the JSON menu configuration file.")
	flag.StringVar(&cfg.SelectedOpt, "r", "", "String coming back from the picker used to find relevant exec command in the config")
	flag.StringVar(&cfg.SelectedOpt, "picker-result", "", "String coming back from the picker used to find relevant exec command in the config")
	flag.StringVar(&cfg.DefaultExec, "e", "", "Command to run on every result.")
	flag.StringVar(&cfg.DefaultExec, "default-exec", "", "Command to run on every result.")

	flag.Parse()

	if cfg.SourceJSONFilePath == "" {
		return nil, fmt.Errorf("-m/--menu-config is mandatory")
	}

	return cfg, nil
}
