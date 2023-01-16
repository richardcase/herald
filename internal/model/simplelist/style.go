package simplelist

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/richardcase/herald/internal/colorscheme"
)

// listStyle is the style of the list.
type listStyle struct {
	colors       colorscheme.Colorscheme
	titleStyle   lipgloss.Style
	noItemsStyle lipgloss.Style
	itemStyle    lipgloss.Style

	bracketStyle lipgloss.Style
	numberStyle  lipgloss.Style
}

// newListStyle creates a new listStyle
func newListStyle(colors colorscheme.Colorscheme) listStyle {
	// Create the new style
	newStyle := listStyle{}

	// Set the colors
	newStyle.colors = colors

	// titleStyle is used to style the title of the list
	newStyle.titleStyle = lipgloss.NewStyle().
		Foreground(colors.Color1).
		MarginLeft(3).
		PaddingBottom(1)

	// newItemsStyle is used to style the message when there are no items
	newStyle.noItemsStyle = lipgloss.NewStyle().
		MarginLeft(3).
		Foreground(colors.Color2).
		Italic(true)

	// newItemsStyle is used to style the items in the list
	newStyle.itemStyle = lipgloss.NewStyle().
		MarginLeft(3).
		Foreground(colors.Color2)

	// bracketStyle is used to style the brackets around the index
	newStyle.bracketStyle = lipgloss.NewStyle().
		Foreground(colors.Color7)

	// numberStyle is used to style the number in the index
	newStyle.numberStyle = lipgloss.NewStyle().
		Foreground(colors.Color6)

	// Return the new style
	return newStyle
}

// styleDescription will style the description of the item
func (s listStyle) styleDescription(description string) string {
	// Create the arrow style
	arrowStyle := lipgloss.NewStyle().
		MarginLeft(10).
		Foreground(s.colors.Color3)

	// Create the text style
	textStyle := lipgloss.NewStyle().
		MarginLeft(1).
		Foreground(s.colors.Color3)

	return arrowStyle.Render("⮡") + textStyle.Render(description)
}

// styleIndex will style the index of the item
func (s listStyle) styleIndex(index int, isSelected bool) string {
	// If the index is the active index render it differently
	numberStyle := s.numberStyle.Copy()
	if isSelected {
		numberStyle = numberStyle.Background(s.colors.Text)
	}

	// Check if the index is a digit
	var indexString string
	if index < 10 {
		// Show a digit
		indexString = fmt.Sprintf("%d", index)
	} else {
		// Show a letter
		indexString = fmt.Sprintf("%c", index+87)
	}
	// Render the whole index
	return lipgloss.NewStyle().
		MarginLeft(3).
		Render(
			s.bracketStyle.Render("[") +
				numberStyle.Render(indexString) +
				s.bracketStyle.Render("]"),
		)
}
