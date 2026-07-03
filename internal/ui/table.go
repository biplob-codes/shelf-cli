package ui

import (
	"fmt"
	"os"

	"github.com/biplob-codes/shelf-cli/internal/store"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

var headerStyle = lipgloss.NewStyle().Bold(true).Padding(0, 1)

const (
	idWidth   = 6
	timeWidth = 22

	minTitleWidth = 10
	maxTitleWidth = 60

	minLinksWidth = 8
	maxLinksWidth = 12

	minURLWidth = 15
	maxURLWidth = 70

	minTagWidth = 8
	maxTagWidth = 28

	fallbackWidth = 100
)

// terminalWidth returns the current terminal width, falling back to a
// sane default when stdout isn't a real terminal (e.g. piped output).
func terminalWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w <= 0 {
		return fallbackWidth
	}
	return w
}

// contentWidth returns the width a column actually needs: the longest of
// its header and its values, clamped to [min, max]. It does NOT try to
// fill available space — that's handled separately, only when the
// terminal is too narrow to fit the natural width.
func contentWidth(header string, values []string, min, max int) int {
	w := len(header)
	for _, v := range values {
		if l := len(v); l > w {
			w = l
		}
	}
	if w < min {
		w = min
	}
	if w > max {
		w = max
	}
	return w
}

func RenderCollectionsTable(collections []store.Collection) string {
	total := terminalWidth()
	overhead := 8 // 4 columns * 2 chars of horizontal padding each

	titles := make([]string, len(collections))
	links := make([]string, len(collections))
	for i, c := range collections {
		titles[i] = c.Title
		links[i] = fmt.Sprintf("%d", c.LinkCount)
	}

	linksWidth := contentWidth("LINKS", links, minLinksWidth, maxLinksWidth)

	// How much room TITLE could take if it needed the whole terminal.
	availableForTitle := total - idWidth - linksWidth - timeWidth - overhead
	if availableForTitle < minTitleWidth {
		availableForTitle = minTitleWidth
	}
	titleCap := maxTitleWidth
	if availableForTitle < titleCap {
		titleCap = availableForTitle
	}
	titleWidth := contentWidth("TITLE", titles, minTitleWidth, titleCap)

	t := table.New().
		BorderTop(false).
		BorderBottom(false).
		BorderLeft(false).
		BorderRight(false).
		BorderColumn(false).
		BorderHeader(false).
		Headers("ID", "TITLE", "LINKS", "CREATED").
		StyleFunc(func(row, col int) lipgloss.Style {
			style := lipgloss.NewStyle().Padding(0, 1)
			if row == table.HeaderRow {
				style = headerStyle
			}
			switch col {
			case 0: // ID
				return style.Width(idWidth)
			case 1: // TITLE
				return style.Width(titleWidth)
			case 2: // LINKS
				return style.Width(linksWidth)
			case 3: // CREATED
				return style.Width(timeWidth)
			default:
				return style
			}
		})

	for _, c := range collections {
		t = t.Row(
			fmt.Sprintf("%d", c.ID),
			truncate(c.Title, titleWidth-2),
			fmt.Sprintf("%d", c.LinkCount),
			FormatTime(c.CreatedAt),
		)
	}

	return t.String()
}

func RenderLinksTable(links []store.Link) string {
	total := terminalWidth()
	overhead := 8 // 4 columns * 2 chars of horizontal padding each

	urls := make([]string, len(links))
	tags := make([]string, len(links))
	for i, l := range links {
		urls[i] = l.URL
		tags[i] = l.Tag
	}

	// Room left for URL + TAG combined if they took the whole terminal.
	availableForBoth := total - idWidth - timeWidth - overhead
	if availableForBoth < minURLWidth+minTagWidth {
		availableForBoth = minURLWidth + minTagWidth
	}

	// TAG first: it's the secondary column, so it only ever grows to
	// fit its own content, capped, and never steals space URL needs.
	tagCap := maxTagWidth
	if availableForBoth-minURLWidth < tagCap {
		tagCap = availableForBoth - minURLWidth
	}
	if tagCap < minTagWidth {
		tagCap = minTagWidth
	}
	tagWidth := contentWidth("TAG", tags, minTagWidth, tagCap)

	// URL gets whatever's left, up to its own natural need.
	urlCap := availableForBoth - tagWidth
	if urlCap > maxURLWidth {
		urlCap = maxURLWidth
	}
	if urlCap < minURLWidth {
		urlCap = minURLWidth
	}
	urlWidth := contentWidth("URL", urls, minURLWidth, urlCap)

	t := table.New().
		BorderTop(false).
		BorderBottom(false).
		BorderLeft(false).
		BorderRight(false).
		BorderColumn(false).
		BorderHeader(false).
		Headers("ID", "URL", "TAG", "ADDED").
		StyleFunc(func(row, col int) lipgloss.Style {
			style := lipgloss.NewStyle().Padding(0, 1)
			if row == table.HeaderRow {
				style = headerStyle
			}
			switch col {
			case 0: // ID
				return style.Width(idWidth)
			case 1: // URL
				return style.Width(urlWidth)
			case 2: // TAG
				return style.Width(tagWidth)
			case 3: // ADDED
				return style.Width(timeWidth)
			default:
				return style
			}
		})

	for _, l := range links {
		t = t.Row(
			fmt.Sprintf("%d", l.ID),
			truncate(l.URL, urlWidth-2),
			truncate(l.Tag, tagWidth-2),
			FormatTime(l.CreatedAt),
		)
	}

	return t.String()
}

func truncate(s string, max int) string {
	if max <= 3 || len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}