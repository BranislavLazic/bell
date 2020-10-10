package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		fmt.Printf("%+v\n", program)
		fmt.Printf("%+v\n", p.Errors)
	}
}
