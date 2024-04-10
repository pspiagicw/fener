package main

import (
	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/handle"
)

var VERSION string = "unversioned"

func main() {
	opts := argparse.Parse(VERSION)

	handle.Handle(opts)
}
