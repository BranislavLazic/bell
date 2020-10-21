package repl

import (
	"bufio"
	"fmt"
	"github.com/branislavlazic/bell/evaluator"
	"io"
	"os"

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
		if len(p.Errors) > 0 {
			parserErrs := fmt.Sprintf("%+v\n", p.Errors)
			_, _ = out.Write([]byte(parserErrs))
		} else {
			evalRes := evaluator.Eval(program)
			evalResult := fmt.Sprintf("%+v\n", evalRes.Inspect())
			_, err := out.Write([]byte(evalResult))
			if err != nil {
				fmt.Println("Failed to write a result.")
				os.Exit(1)
			}
		}
	}
}
