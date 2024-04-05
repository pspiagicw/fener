package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
	"github.com/sanity-io/litter"
)

func main() {
	for true {

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf(">>> ")
		value, err := reader.ReadString('\n')
		if err != nil {
			goreland.LogFatal("Error reading string: %v", err)
		}

		l := lexer.New(value)

		p := parser.New(l)

		program := p.Parse()

		fmt.Println(litter.Sdump(program))
		fmt.Println(program)

		for _, err := range p.Errors() {
			fmt.Println(err)
		}

	}
}
