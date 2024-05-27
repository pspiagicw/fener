package repl

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/eval"
	"github.com/pspiagicw/fener/help"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
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

	rl := initREPL()
	defer rl.Close()

	for true {

		line := getInput(rl)

		ast, errors := parseLine(line)

		for _, err := range errors {
			goreland.LogError(err)
		}

		if opts.PrintAST {
			printAST(ast)
		}

		e := eval.New(func(err error) {
			goreland.LogError(err.Error())
		})

		fmt.Println(e.Eval(ast))
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

func getInput(r *readline.Instance) string {
	line, err := r.Readline()
	if err != nil {
		goreland.LogFatal("Error reading input from prompt: %v", err)
	}
	line = strings.TrimSpace(line)

	return line
}

func initREPL() *readline.Instance {
	r, err := readline.NewEx(&readline.Config{
		Prompt:          ">>> ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "^D",
	})

	if err != nil {
		goreland.LogFatal("Error initalizing readline: %v", err)
	}

	return r
}
