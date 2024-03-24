package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pspiagicw/fener/lexer"
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

		l := lexer.NewLexer(value)

		t := l.Next()

		for t.Type != "EOF" {
			if l.Error() != nil {
				goreland.LogError("Error: %v", l.Error())
			}
			fmt.Println(t)
			t = l.Next()
		}

	}
}
