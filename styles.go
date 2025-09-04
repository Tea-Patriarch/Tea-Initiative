package main

import (
	"github.com/charmbracelet/lipgloss"
)

var selected = lipgloss.NewStyle().Background(lipgloss.Color("9")).Foreground(lipgloss.Color("0"))
var line = lipgloss.NewStyle().Background(lipgloss.Color("235"))
var wait = lipgloss.NewStyle().Foreground(lipgloss.Color("011"))
var rowStyle = lipgloss.NewStyle().Padding(0, 2)
var headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15"))

var idWidth, nameWidth, initWidth = 2, 20, 15

var box = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	BorderForeground(lipgloss.Color("15"))
