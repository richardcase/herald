package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/richardcase/herald/internal/ui/context"
)

func GetNotificationTextStyle(
	ctx *context.ProgramContext,
	unread bool,
) lipgloss.Style {
	if unread {
		return lipgloss.NewStyle().Foreground(ctx.Theme.PrimaryText)
	}
	return lipgloss.NewStyle().Foreground(ctx.Theme.FaintText)
}

func RenderNotificationTitle(
	ctx *context.ProgramContext,
	unread bool,
	title string,
) string {
	rTitle := GetNotificationTextStyle(ctx, unread).Render(title)

	res := fmt.Sprintf("%s", rTitle)
	return res
}
