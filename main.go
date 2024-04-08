package main

import (
	"github.com/pspiagicw/fener/argparse"
	"github.com/pspiagicw/fener/handle"
)

var VERSION string = "unversioned"

func main() {
	opts := argparse.Parse(VERSION)

	handle.Handle(opts)

	// for true {
	//
	// 	reader := bufio.NewReader(os.Stdin)
	//
	// 	fmt.Printf(">>> ")
	// 	value, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		goreland.LogFatal("Error reading string: %v", err)
	// 	}
	//
	// 	l := lexer.New(value)
	//
	// 	p := parser.New(l)
	//
	// 	program := p.Parse()
	//
	// 	fmt.Println(litter.Sdump(program))
	// 	fmt.Println(program)
	//
	// 	for _, err := range p.Errors() {
	// 		fmt.Println(err)
	// 	}
	//
	// }
}
