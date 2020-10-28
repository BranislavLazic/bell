package repl

import (
	"fmt"
	"github.com/branislavlazic/bell/evaluator"
	"github.com/branislavlazic/bell/lexer"
	"github.com/branislavlazic/bell/object"
	"github.com/branislavlazic/bell/parser"
	input "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"
	"log"
)

const textPrompt = ">> "
const inputForegroundColor = "147"
const outputForegroundColor = "42"

type model struct {
	textInput   input.Model
	result      string
	env         *object.Environment
	expressions []string
	isQuit      bool
}

func initialModel() model {
	inputModel := input.NewModel()
	inputModel.Prompt = textPrompt
	inputModel.TextColor = inputForegroundColor
	inputModel.Focus()

	return model{
		textInput:   inputModel,
		result:      "",
		env:         object.NewEnvironment(),
		expressions: []string{},
		isQuit:      false,
	}
}

func (m model) Init() tea.Cmd {
	return input.Blink(m.textInput)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			textInputValue := m.textInput.Value()
			if !m.execCommand(textInputValue) {
				m.expressions = append(m.expressions, textInputValue)
				l := lexer.New(textInputValue)
				p := parser.New(l)
				program := p.ParseProgram()
				if len(p.Errors) > 0 {
					for _, err := range p.Errors {
						m.result = outputEvalResult(m.result, textInputValue, err)
					}
				} else {
					evalRes := evaluator.Eval(program, m.env)
					m.result = outputEvalResult(m.result, textInputValue, evalRes.Inspect())
				}
			}
			if m.isQuit {
				return m, tea.Quit
			}
			m.textInput.Reset()
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	m.textInput, cmd = input.Update(msg, m.textInput)
	return m, cmd
}

func (m model) View() string {
	color := te.ColorProfile().Color
	output := fmt.Sprintf(
		"%s\n%s\n",
		m.result,
		input.View(m.textInput),
	)
	return te.String(output).Foreground(color(outputForegroundColor)).String()
}

func outputEvalResult(currentRes string, textInputValue string, evalRes string) string {
	return fmt.Sprintf("%s\n>> %s\n%s\n", currentRes, textInputValue, evalRes)
}

func (m *model) execCommand(cmd string) bool {
	switch cmd {
	case ":help":
		m.result = `
:clear - clear the console output
:quit - (or Ctrl+C) quit the session
`
		return true
	case ":clear":
		m.result = ""
		m.env = object.NewEnvironment()
		m.expressions = []string{}
		return true
	case ":quit":
		m.isQuit = true
		return true
	}
	return false
}

func Start() {
	fmt.Println("(Type :help to show of commands)")
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
