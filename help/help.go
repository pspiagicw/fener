package help

import (
	"github.com/pspiagicw/pelp"
)

func Usage() {
	pelp.Print("A interpreted language")

	pelp.Aligned(
		"commands",
		[]string{"help", "run", "repl", "version"},
		[]string{"Show this help message", "Run a file", "Start a repl", "Show version"},
	)
}
func Version(version string) {
	pelp.Version("fener", version)
}
func Handle(args []string, version string) {
	if len(args) == 0 {
		Usage()
		return
	}

	cmd := args[0]

	switch cmd {
	case "version":
		Version(version)
	case "repl":
		Repl()
	case "run":
		Run()
	}
}

func Run() {
	pelp.Print("Run a file")

	pelp.Flags("flags", []string{"print-ast"}, []string{"Print the AST of the program"})
}

func Repl() {
	pelp.Print("Start fener repl")

	pelp.Flags("flags", []string{"print-ast"}, []string{"Print the AST of the program"})
}
