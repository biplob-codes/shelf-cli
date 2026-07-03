package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorSuccess = lipgloss.Color("42")
	colorError   = lipgloss.Color("196")
	colorMuted   = lipgloss.Color("240")
	colorAccent  = lipgloss.Color("39")
)

var (
	Success = lipgloss.NewStyle().Foreground(colorSuccess)
	Error   = lipgloss.NewStyle().Foreground(colorError)
	Muted   = lipgloss.NewStyle().Foreground(colorMuted)
	Accent  = lipgloss.NewStyle().Foreground(colorAccent)
)
