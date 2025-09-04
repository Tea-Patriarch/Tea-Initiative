package main

import (
	"os"
	_ "strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// defining model
func initialModel() model {
	return model{}
}

// inicialising model
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		if m.inputMode {
			return m.InputFun(msg)
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", "space", "enter":
			if m.cursor < len(m.characters)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.characters) - 1
			}
		case "i":
			m.inputMode = true
			m.inputBuffer = ""
		case "w":
			if len(m.characters) != 0 {
				m.characters[m.cursor].isVating = !m.characters[m.cursor].isVating
			}
		case "d":
			if len(m.characters) != 0 {
				m.characters = append(m.characters[:m.cursor], m.characters[m.cursor+1:]...)
				if m.cursor > 0 {
					m.cursor--
				}
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	var s string

	s += box.Padding(0, 0).Render("Initiative") + "\n"

	s += box.Render(
		MakeTable(m),
		"\n\n",
		// strconv.Itoa(m.cursor),
	)
	s += MakeBottom(m)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		s,
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		println("Error:", err)
		os.Exit(1)
	}
}
