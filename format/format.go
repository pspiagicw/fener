package format

import (
	"fmt"
	"os"
	"strings"

	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/ast"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
)

func Handle(opts *argparse.Opts) {

	for _, arg := range opts.Args {
		// Format the file

		ast, errors := parseFile(arg)

		if len(errors) > 0 {
			for _, err := range errors {
				goreland.LogError(err)
			}
			goreland.LogFatal("Parsing failed!!!")
		}

		output := format(ast)

		fmt.Println(output)
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
func format(program *ast.Program) string {
	var out strings.Builder
	for _, s := range program.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}
