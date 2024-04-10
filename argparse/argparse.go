package argparse

import "flag"
import "github.com/pspiagicw/fener/help"

type Opts struct {
	Version string
	Args    []string

	// REPL
	PrintAST bool
}

func Parse(version string) *Opts {
	o := &Opts{}

	o.Version = version
	flag.Usage = help.Usage

	flag.Parse()

	o.Args = flag.Args()
	return o
}
