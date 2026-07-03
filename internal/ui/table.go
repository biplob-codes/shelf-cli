package ui

import (
	"fmt"

	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func RenderCollectionsTable(collections []store.Collection) string {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(Muted).
		Headers("ID", "TITLE", "LINKS", "CREATED").
		StyleFunc(func(row, col int) lipgloss.Style {
			switch col {
			case 0: // ID
				return lipgloss.NewStyle().Width(6).Padding(0, 1)
			case 3: // CREATED
				return lipgloss.NewStyle().Width(22).Padding(0, 1)
			default:
				return lipgloss.NewStyle().Padding(0, 1)
			}
		})

	for _, c := range collections {
		t = t.Row(
			fmt.Sprintf("%d", c.ID),
			c.Title,
			fmt.Sprintf("%d", c.LinkCount),
			FormatTime(c.CreatedAt),
		)
	}
	return t.String()
}

func RenderLinksTable(links []store.Link) string {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(Muted).
		Headers("ID", "URL", "TAG", "ADDED").
		StyleFunc(func(row, col int) lipgloss.Style {
			switch col {
			case 0: // ID
				return lipgloss.NewStyle().Width(6).Padding(0, 1)
			case 2: // TAG
				return lipgloss.NewStyle().Width(20).Padding(0, 1)
			case 3: // ADDED
				return lipgloss.NewStyle().Width(22).Padding(0, 1)
			default:
				return lipgloss.NewStyle().Padding(0, 1)
			}
		})

	for _, l := range links {
		t = t.Row(
			fmt.Sprintf("%d", l.ID),
			truncate(l.URL, 60),
			l.Tag,
			FormatTime(l.CreatedAt),
		)
	}
	return t.String()
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}