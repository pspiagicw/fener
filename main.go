package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/pspiagicw/fener/lexer"
	"github.com/pspiagicw/fener/parser"
	"github.com/pspiagicw/goreland"
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

		for _, err := range p.Errors() {
			fmt.Println(err)
		}

		spew.Config.DisablePointerAddresses = true
		spew.Config.Indent = "\t"

		spew.Printf("%#v\n", program)

	}
}
