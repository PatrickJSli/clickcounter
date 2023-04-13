package bubble

import (
	"fmt"

	"github.com/PatrickJSli/clickcounter/xinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	count  int
	width  int
	height int
}

type clickMsg int

var style = lipgloss.NewStyle().
	Padding(1, 1).
	Align(lipgloss.Center, lipgloss.Center)

func (m model) getClick() tea.Msg {
	xinput.WaitForClick()
	return clickMsg(1)
}

func (m model) Init() tea.Cmd {
	return m.getClick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		style.Padding(m.height/2, m.width/2)
	case clickMsg:
		m.count += 1
		return m, m.getClick
	}
	return m, nil
}

func (m model) View() string {
	return style.Render(fmt.Sprintf("%d", m.count))
}

func Run() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
