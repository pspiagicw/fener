package handle

import (
	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/format"
	"github.com/pspiagicw/fener/help"
	"github.com/pspiagicw/fener/repl"
	"github.com/pspiagicw/fener/run"
	"github.com/pspiagicw/goreland"
)

var handles = map[string]func(*argparse.Opts){
	"repl": repl.Handle,
	"run":  run.Handle,
	"version": func(opts *argparse.Opts) {
		help.Version(opts.Version)
	},
	"help": func(opts *argparse.Opts) {
		help.Handle(opts.Args, opts.Version)
	},
	"format": format.Handle,
}

func Handle(opts *argparse.Opts) {
	if len(opts.Args) == 0 {
		help.Usage()
		return
	}

	cmd := opts.Args[0]
	// Remove the command from the arguments
	opts.Args = opts.Args[1:]

	if h, ok := handles[cmd]; ok {
		h(opts)
	} else {
		goreland.LogFatal("Unknown command: %s", cmd)
	}

}
