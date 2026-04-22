package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var loginHeader string = `Test 








		**End of Test**
	`

type Model struct {
	viewport  viewport.Model
	textInput textinput.Model
	log       []string
	width     int
	height    int
}

func InitialModel(width, height int) Model {
	m := Model{
		width:  width,
		height: height,
	}

	ti := textinput.New()
	ti.Placeholder = "> "
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	m.textInput = ti
	m.log = append(m.log, "> "+loginHeader)

	m.viewport = viewport.New(width, height/2)
	m.viewport.MouseWheelEnabled = true
	m.viewport.SetContent(strings.Join(m.log, "\n"))

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.log = append(m.log, "> "+m.textInput.Value())
			m.textInput.Reset()
			content := strings.Join(m.log, "\n")
			m.viewport.SetContent(content)
			m.viewport.GotoBottom()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Height = m.height - 15
		m.viewport.Width = m.width - 2
	}

	m.textInput, cmd = m.textInput.Update(msg)
	m.viewport, cmd = m.viewport.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	//content := strings.Join(m.log, "\n")
	//m.viewport.SetContent(content)
	//m.viewport.GotoBottom()

	var style = lipgloss.NewStyle().
		BorderStyle(lipgloss.ASCIIBorder()).
		Width(m.width - 2)
	vPortBorder := style.Render(m.viewport.View())
	return fmt.Sprintf("%s\n\n%s",
		vPortBorder,
		m.textInput.View(),
	)
}
