package main

import (
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

func MakeTable(m model) string {
	var rows []string

	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		rowStyle.Width(idWidth).Render("ID"),
		rowStyle.Width(nameWidth).Render("Name"),
		rowStyle.Width(initWidth).Render("Iniciative"),
	)
	header = line.Render(header)
	rows = append(rows, headerStyle.Render(header))

	for i, c := range m.characters {
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			rowStyle.Width(idWidth).Render(" "+strconv.Itoa(i)),
			rowStyle.Width(nameWidth).Render(c.name),
			rowStyle.Width(initWidth-1).Render(strconv.Itoa(c.initiative)),
		)
		if i%2 == 1 && i != m.cursor {
			row = line.Render(row)
		}
		if m.characters[i].isVating == true {
			row += wait.Render("")
		}
		if i == m.cursor {
			row = selected.Render(row)
		}
		rows = append(rows, row)
	}

	table := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return table
}

func MakeBottom(m model) string {
	var s string
	if m.inputMode {
		s += "\n\n" + line.Render(wait.Render("> "+m.inputBuffer+" <"))
	} else {
		s += "\n\n" + lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("011")).Render("i") + " to add character    " +
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("011")).Render("j/k") + " to move    " +
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("011")).Render("q") + " to quit"
		s += "\n" + lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("011")).Render("d") + " to delite character    " +
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("011")).Render("w") + " to mark"
	}

	return s
}
