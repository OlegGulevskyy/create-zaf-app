package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type promptInputProps struct {
	placeholder string
	title       string
}

func (p *Project) promptInput(props promptInputProps) {
	initialModel := initialModel(props.placeholder)
	initialModel.project = p
	initialModel.title = props.title

	prog := tea.NewProgram(initialModel)
	if _, err := prog.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type modelInput struct {
	textInput textinput.Model
	title     string
	err       error
	project   *Project
}

func initialModel(placeholder string) modelInput {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return modelInput{
		textInput: ti,
		err:       nil,
	}
}

func (m modelInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m modelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	m.project.selectedInputItem = m.textInput.Value()
	return m, cmd
}

func (m modelInput) View() string {
	return fmt.Sprintf(
		"%s?\n\n%s\n\n%s",
		m.title,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
