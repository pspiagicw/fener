package repl

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
	"github.com/sanity-io/litter"
)

func parseReplArgs(opts *argparse.Opts) {
	flag := flag.NewFlagSet("fener repl", flag.ExitOnError)

	flag.BoolVar(&opts.PrintAST, "print-ast", false, "Print the AST of the program")

	flag.Parse(opts.Args)
}

func Handle(opts *argparse.Opts) {

	parseReplArgs(opts)

	reader := bufio.NewReader(os.Stdin)

	for true {
		line := getLine(reader)

		ast, errors := parseLine(line)

		for _, err := range errors {
			goreland.LogError(err)
		}

		if opts.PrintAST {
			printAST(ast)
		}
	}
}
func parseLine(line string) (*ast.Program, []string) {
	l := lexer.New(line)
	p := parser.New(l)
	program := p.Parse()

	return program, p.Errors()

}

func getLine(rw *bufio.Reader) string {

	fmt.Printf(">>> ")

	value, err := rw.ReadString('\n')
	if err != nil {
		goreland.LogFatal("Error reading string: %v", err)
	}
	return value
}
func printAST(program *ast.Program) {
	litter.Dump(program)
}
