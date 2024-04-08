package argparse

import "flag"

type Opts struct {
	Version string
	Args    []string

	// REPL
	PrintAST bool
}

func Parse(version string) *Opts {
	o := &Opts{}

	o.Version = version

	flag.Parse()
	o.Args = flag.Args()
	return o
}
