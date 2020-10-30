package main

import (
	"fmt"
	"github.com/branislavlazic/bell/evaluator"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/object"
	"github.com/branislavlazic/bell/parser"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func LoadBellFile(fileName string) ([]byte, error) {
	if !strings.HasSuffix(fileName, ".bell") {
		log.Fatalf("Invalid file name. extension must be bell.")
	}
	return ioutil.ReadFile(fileName)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Provide a source code file with .bell extension")
	}
	file, err := LoadBellFile(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot read a file.")
	}
	env := object.NewEnvironment()
	l := lexer.New(string(file))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors) > 0 {
		for _, err := range p.Errors {
			fmt.Println(err)
		}
	} else {
		evaluator.Eval(program, env)
		for _, obj := range evaluator.SysOut {
			fmt.Println(obj.Inspect())
		}
	}
}
