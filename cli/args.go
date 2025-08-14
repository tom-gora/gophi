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
	ExecCmd            string
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
		tbl.AddRow("-e, --exec", "string", "Command to execute as returned from picker.", "")
		tbl.Print()
	}

	flag.StringVar(&cfg.SourceJSONFilePath, "m", "", "Path to the JSON menu configuration file.")
	flag.StringVar(&cfg.SourceJSONFilePath, "menu-config", "", "Path to the JSON menu configuration file.")
	flag.StringVar(&cfg.ExecCmd, "e", "", "Command to execute as returned from picker.")
	flag.StringVar(&cfg.ExecCmd, "exec", "", "Command to execute as returned from picker.")

	flag.Parse()

	if cfg.SourceJSONFilePath == "" {
		return nil, fmt.Errorf("-m/--menu-config is mandatory")
	}

	return cfg, nil
}
