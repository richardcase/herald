package feed

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/richardcase/herald/internal/colorscheme"
)

// style is the style of the feed tab.
type style struct {
	columnStyle  lipgloss.Style
	focusedStyle lipgloss.Style
}

// newStyle creates a new style for the feed tab.
func newStyle(colors colorscheme.Colorscheme) style {
	return style{
		columnStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(colors.TextDark),
		focusedStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(colors.Color1),
	}
}
