package context

import (
	bbHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"

	"github.com/richardcase/herald/internal/ui/common"
	"github.com/richardcase/herald/internal/ui/theme"
)

type Styles struct {
	Colors struct {
		NewNotification    lipgloss.AdaptiveColor
		UnreadNotification lipgloss.AdaptiveColor
		ReadNotification   lipgloss.AdaptiveColor
	}

	Common common.CommonStyles

	NotificationSidebar struct {
		PillStyle lipgloss.Style
	}

	Help struct {
		Text         lipgloss.Style
		KeyText      lipgloss.Style
		BubbleStyles bbHelp.Styles
	}
	Pager struct {
		Height int
		Root   lipgloss.Style
	}
	Section struct {
		ContainerPadding int
		ContainerStyle   lipgloss.Style
		SpinnerStyle     lipgloss.Style
		EmptyStateStyle  lipgloss.Style
		KeyStyle         lipgloss.Style
	}
	NotificationSection struct {
		CiCellWidth        int
		LinesCellWidth     int
		UpdatedAtCellWidth int
		RepoCellWidth      int
		AuthorCellWidth    int
	}
	Sidebar struct {
		BorderWidth    int
		PagerHeight    int
		ContentPadding int
		Root           lipgloss.Style
		PagerStyle     lipgloss.Style
	}
	ListViewPort struct {
		PagerStyle lipgloss.Style
	}
	Table struct {
		CellStyle                lipgloss.Style
		SelectedCellStyle        lipgloss.Style
		TitleCellStyle           lipgloss.Style
		SingleRuneTitleCellStyle lipgloss.Style
		HeaderStyle              lipgloss.Style
		RowStyle                 lipgloss.Style
	}
	Tabs struct {
		Tab            lipgloss.Style
		ActiveTab      lipgloss.Style
		TabSeparator   lipgloss.Style
		TabsRow        lipgloss.Style
		ViewSwitcher   lipgloss.Style
		ActiveView     lipgloss.Style
		ViewsSeparator lipgloss.Style
		InactiveView   lipgloss.Style
	}
}

func InitStyles(theme theme.Theme) Styles {
	var s Styles

	s.Colors.NewNotification = lipgloss.AdaptiveColor{
		Light: "#42A0FA",
		Dark:  "#42A0FA",
	}
	s.Colors.ReadNotification = lipgloss.AdaptiveColor{
		Light: "#C38080",
		Dark:  "#C38080",
	}
	s.Colors.UnreadNotification = lipgloss.AdaptiveColor{
		Light: "#3DF294",
		Dark:  "#3DF294",
	}

	s.Common = common.BuildStyles(theme)

	s.NotificationSidebar.PillStyle = s.Common.MainTextStyle.Copy().
		Foreground(theme.InvertedText).
		PaddingLeft(1).
		PaddingRight(1)

	s.Help.Text = lipgloss.NewStyle().Foreground(theme.SecondaryText)
	s.Help.KeyText = lipgloss.NewStyle().Foreground(theme.PrimaryText)
	s.Help.BubbleStyles = bbHelp.Styles{
		ShortDesc:      s.Help.Text.Copy().Foreground(theme.FaintText),
		FullDesc:       s.Help.Text.Copy().Foreground(theme.FaintText),
		ShortSeparator: s.Help.Text.Copy().Foreground(theme.SecondaryBorder),
		FullSeparator:  s.Help.Text.Copy(),
		FullKey:        s.Help.KeyText.Copy(),
		ShortKey:       s.Help.KeyText.Copy(),
		Ellipsis:       s.Help.Text.Copy(),
	}

	//s.CommentBox.Text = s.Help.Text.Copy()

	s.Pager.Height = 2
	s.Pager.Root = lipgloss.NewStyle().
		Height(s.Pager.Height).
		MaxHeight(s.Pager.Height).
		PaddingTop(1).
		Bold(true).
		Foreground(theme.FaintText)

	s.Section.ContainerPadding = 1
	s.Section.ContainerStyle = lipgloss.NewStyle().
		Padding(0, s.Section.ContainerPadding)
	s.Section.SpinnerStyle = lipgloss.NewStyle().Padding(0, 1)
	s.Section.EmptyStateStyle = lipgloss.NewStyle().
		Faint(true).
		PaddingLeft(1).
		MarginBottom(1)
	s.Section.KeyStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Background(theme.SelectedBackground).
		Padding(0, 1)

	s.NotificationSection.CiCellWidth = lipgloss.Width(" CI ")
	s.NotificationSection.LinesCellWidth = lipgloss.Width(" 123450 / -123450 ")
	s.NotificationSection.UpdatedAtCellWidth = lipgloss.Width("2mo ago")
	s.NotificationSection.RepoCellWidth = 15
	s.NotificationSection.AuthorCellWidth = 15

	s.Sidebar.BorderWidth = 1
	s.Sidebar.ContentPadding = 2
	s.Sidebar.Root = lipgloss.NewStyle().
		Padding(0, s.Sidebar.ContentPadding).
		BorderLeft(true).
		BorderStyle(lipgloss.Border{
			Top:         "",
			Bottom:      "",
			Left:        "│",
			Right:       "",
			TopLeft:     "",
			TopRight:    "",
			BottomRight: "",
			BottomLeft:  "",
		}).
		BorderForeground(theme.PrimaryBorder)
	s.Sidebar.PagerStyle = lipgloss.NewStyle().
		Height(s.Sidebar.PagerHeight).
		Bold(true).
		Foreground(theme.FaintText)

	s.ListViewPort.PagerStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Background(theme.SelectedBackground).
		Foreground(theme.FaintText)

	s.Table.CellStyle = lipgloss.NewStyle().PaddingLeft(1).
		PaddingRight(1).
		MaxHeight(1)
	s.Table.SelectedCellStyle = s.Table.CellStyle.Copy().
		Background(theme.SelectedBackground)
	s.Table.TitleCellStyle = s.Table.CellStyle.Copy().
		Bold(true).
		Foreground(theme.PrimaryText)
	s.Table.SingleRuneTitleCellStyle = s.Table.TitleCellStyle.Copy().
		Width(common.SingleRuneWidth)
	s.Table.HeaderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(theme.FaintBorder).
		BorderBottom(true)
	s.Table.RowStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(theme.FaintBorder)

	s.Tabs.Tab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 2)
	s.Tabs.ActiveTab = s.Tabs.Tab.
		Copy().
		Faint(false).
		Bold(true).
		Background(theme.SelectedBackground).
		Foreground(theme.PrimaryText)
	s.Tabs.TabSeparator = lipgloss.NewStyle().
		Foreground(theme.SecondaryBorder)
	s.Tabs.TabsRow = lipgloss.NewStyle().
		Height(common.TabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottomForeground(theme.PrimaryBorder)
	s.Tabs.ViewSwitcher = lipgloss.NewStyle().
		Background(theme.SecondaryText).
		Foreground(theme.InvertedText).
		Padding(0, 1).
		Bold(true)

	s.Tabs.ActiveView = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Bold(true).
		Background(theme.SelectedBackground)
	s.Tabs.ViewsSeparator = lipgloss.NewStyle().
		BorderForeground(theme.PrimaryBorder).
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true)
	s.Tabs.InactiveView = lipgloss.NewStyle().
		Background(theme.FaintBorder).
		Foreground(theme.SecondaryText)

	return s
}
