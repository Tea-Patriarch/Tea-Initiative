package main

import (
	"sort"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) InputFun(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			parts := strings.Split(m.inputBuffer, " ")
			if len(parts) >= 2 {
				name := parts[0]
				initiative, err := strconv.Atoi(parts[1])
				if err == nil {
					m.characters = append(m.characters, chara{name: name, initiative: initiative})
				}
			}
			m.inputBuffer = ""
			m.inputMode = false
			m.characters = sortByInitiative(m.characters)

		case tea.KeyBackspace:
			if len(m.inputBuffer) > 0 {
				m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
			}

		default:
			m.inputBuffer += msg.String()
		}
	}
	return m, nil
}
func sortByInitiative(chars []chara) []chara {
	sort.Slice(chars, func(i, j int) bool {
		return chars[i].initiative < chars[j].initiative
	})

	return chars
}
