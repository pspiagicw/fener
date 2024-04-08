package run

import (
	"flag"
	"os"

	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
	"github.com/sanity-io/litter"
)

func parseRunArgs(opts *argparse.Opts) {

	flag := flag.NewFlagSet("fener run", flag.ExitOnError)

	flag.BoolVar(&opts.PrintAST, "print-ast", false, "Print the AST of the program")

	flag.Parse(opts.Args)

	opts.Args = flag.Args()
}

func Handle(opts *argparse.Opts) {

	parseRunArgs(opts)

	for _, arg := range opts.Args {
		ast, errors := parseFile(arg)

		if len(errors) > 0 {
			for _, err := range errors {
				goreland.LogError(err)
			}
			goreland.LogFatal("Parsing failed!!!")
		}

		if opts.PrintAST {
			printAST(ast)
		}
	}
}
func parseFile(filename string) (*ast.Program, []string) {
	contents, err := os.ReadFile(filename)

	if err != nil {
		goreland.LogFatal("Error reading file: %v", err)
	}

	l := lexer.New(string(contents))
	p := parser.New(l)
	program := p.Parse()

	return program, p.Errors()

}
func printAST(program *ast.Program) {
	litter.Dump(program)
}
