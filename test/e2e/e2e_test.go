package e2e

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/branislavlazic/bell/evaluator"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/object"
	"github.com/branislavlazic/bell/parser"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var godogs int

var evalResult string
var parserErrors []string

func program(prog *godog.DocString) error {
	env := object.NewEnvironment()
	l := lexer.New(prog.Content)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors) > 0 {
		parserErrors = p.Errors
	} else {
		evalRes := evaluator.Eval(program, env)
		evalResult = evalRes.Inspect()
	}
	return nil
}

func resultIs(res *godog.DocString) error {
	if res.Content != evalResult {
		return fmt.Errorf("incorrect evaluation result. expected=%s, got=%s", res.Content, evalResult)
	}
	return nil
}

func errorIs(res *godog.DocString) error {
	if res.Content != parserErrors[0] {
		return fmt.Errorf("incorrect result. expected=%s, got=%s", res.Content, parserErrors[0])
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		godogs = 0
		evalResult = ""
		parserErrors = []string{}
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		godogs = 0
		evalResult = ""
		parserErrors = []string{}
	})
	ctx.Step(`^the program$`, program)
	ctx.Step(`^the result is$`, resultIs)
	ctx.Step(`^the error is$`, errorIs)
}

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
	Tags:   "@wip",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()
	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
