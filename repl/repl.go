package repl

import (
	"bufio"
	"fmt"
	"github.com/branislavlazic/bell/evaluator"
	"github.com/branislavlazic/bell/object"
	"io"
	"os"

	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
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
			for _, err := range p.Errors {
				parserErrs := fmt.Sprintf("%s\n", err)
				_, _ = out.Write([]byte(parserErrs))
			}
		} else {
			evalRes := evaluator.Eval(program, env)
			evalResult := fmt.Sprintf("%+v\n", evalRes.Inspect())
			_, err := out.Write([]byte(evalResult))
			if err != nil {
				fmt.Println("Failed to write a result.")
				os.Exit(1)
			}
		}
	}
}
