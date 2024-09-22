package main

import (
	"github.com/charmbracelet/lipgloss"
)

var red = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("9"))

var green = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("86"))

var (
	FAIL = red.Render("FAIL")
	PASS = green.Render("PASS")
)
