package main

import (
	"fmt"
	"os"

	"github.com/branislavlazic/bell/repl"
)

func main() {
	fmt.Println("Welcome to Bell REPL!")
	repl.Start(os.Stdin, os.Stdout)
}
