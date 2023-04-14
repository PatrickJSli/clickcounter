package bubble

import (
	"fmt"
	"time"

	"github.com/PatrickJSli/clickcounter/xinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	figure "github.com/common-nighthawk/go-figure"
)

type model struct {
	count           int
	timedCount      int
	width           int
	height          int
	startTime       time.Time
	clicksPerMinute float64
}

var style = lipgloss.NewStyle().
	Background(lipgloss.Color("#232323")).
	Foreground(lipgloss.Color("#00b894")).
	Padding(1, 1).
	Align(lipgloss.Center, lipgloss.Center)

var helpStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#232323")).
	Foreground(lipgloss.Color("#00b894"))

func newModel() model {
	return model{
		startTime: time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return xinput.WaitForClick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			xinput.QuitXinput()
			return m, tea.Quit
		case "r":
			m.timedCount = 0
			m.startTime = time.Now()
			m.clicksPerMinute = 0
		case "R":
			m.count = 0
			m.timedCount = 0
			m.startTime = time.Now()
			m.clicksPerMinute = 0
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case xinput.ClickMsg:
		m.count += 1
		m.timedCount += 1
		m.clicksPerMinute = float64(m.timedCount) / time.Since(m.startTime).Minutes()
		return m, xinput.WaitForClick
	case xinput.CommandError:
		panic("xinput command failed")
		//return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	text := figure.NewFigure(fmt.Sprintf("%d", m.count), "roman", true).String()
	text += fmt.Sprintf("Clicks per minute: %.2f", m.clicksPerMinute)
	helpText := "r Reset clicks per minute • R Reset all • q quit"
	width, height := lipgloss.Size(text)
	style.Padding((m.height-1)/2-height/2, m.width/2-width/2)
	return lipgloss.JoinVertical(lipgloss.Top, style.Render(text), helpStyle.Render(helpText))
}

func Run() {
	p := tea.NewProgram(newModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
