package repl

import (
	"flag"
	"fmt"

	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/eval"
	"github.com/pspiagicw/fener/help"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/object"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/regolith"
	"github.com/sanity-io/litter"
)

func parseReplArgs(opts *argparse.Opts) {
	flag := flag.NewFlagSet("fener repl", flag.ExitOnError)

	flag.Usage = help.Repl

	flag.BoolVar(&opts.PrintAST, "print-ast", false, "Print the AST of the program")

	flag.Parse(opts.Args)
}

func Handle(opts *argparse.Opts) {

	parseReplArgs(opts)

	rg, err := regolith.New(&regolith.Config{
		StartWords: []string{"if", "fn", "while", "class"},
		EndWords:   []string{"end"},
	})

	if err != nil {
		goreland.LogFatal("Error initializing regolith: %v", err)
	}

	defer rg.Close()

	env := object.NewEnvironment()

	for true {

		line, err := rg.Input()

		if err != nil {
			goreland.LogFatal("Error reading input: %v", err)
		}

		ast, errors := parseLine(line)

		if len(errors) > 0 {
			for _, err := range errors {
				goreland.LogError(err)
			}
			continue
		}

		if opts.PrintAST {
			printAST(ast)
		}

		e := eval.New(func(err error) {
			goreland.LogError(err.Error())
		})

		fmt.Println(e.Eval(ast, env))
	}
}
func parseLine(line string) (*ast.Program, []string) {
	l := lexer.New(line)
	p := parser.New(l)
	program := p.Parse()

	return program, p.Errors()

}

func printAST(program *ast.Program) {
	litter.Dump(program)
}
