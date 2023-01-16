package welcome

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/richardcase/herald/internal/backend"
	"github.com/richardcase/herald/internal/colorscheme"
	"github.com/richardcase/herald/internal/model/simplelist"
	"github.com/richardcase/herald/internal/model/tab"
)

type Model struct {
	colors colorscheme.Colorscheme
	width  int
	height int
	title  string
	loaded bool
	list   simplelist.Model

	// reader is a function which returns a tea.Cmd which will be executed
	// when the tab is initialized
	reader func() tea.Cmd
}

// New creates a new welcome tab with sensible defaults
func New(colors colorscheme.Colorscheme, width, height int, title string, reader func() tea.Cmd) Model {
	return Model{
		colors: colors,
		width:  width,
		height: height,
		title:  title,
		reader: reader,
	}
}

// Title returns the title of the tab
func (m Model) Title() string {
	return m.title
}

// Type returns the type of the tab
func (m Model) Type() tab.Type {
	return tab.Welcome
}

// Help returns the help for the tab
func (m Model) Help() tab.Help {
	return tab.Help{
		tab.KeyBind{Key: "enter", Description: "Open"},
		tab.KeyBind{Key: "ctrl+n", Description: "New"},
		//tab.KeyBind{Key: "ctrl+e", Description: "Edit"},
		//tab.KeyBind{Key: "ctrl+d", Description: "Delete"},
	}
}

// Init initializes the tab
func (m Model) Init() tea.Cmd {
	return m.reader()
}

// Update the variables of the tab
func (m Model) Update(msg tea.Msg) (tab.Tab, tea.Cmd) {
	// Wait for items to be loaded
	if !m.loaded {
		_, ok := msg.(backend.FetchSuccessMessage)
		if !ok {
			return m, nil
		}

		// Initialize the list of categories, items will be set later
		m.list = simplelist.New(m.colors, "Categories", m.height, true)

		// Add the categories
		m.loaded = true
	}

	switch msg := msg.(type) {
	case backend.FetchSuccessMessage:
		// Update the list of categories
		m.list.SetItems(msg.Items)

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			//TODO
			// Add a new tab with the selected category
			if !m.list.IsEmpty() {
				return m, tab.NewTab(m.list.SelectedItem().FilterValue(), tab.Filter)
			}

			// If the list is empty, return nothing
			return m, nil

		case "ctrl+n":
			// Add a new category
			//TODO
			//return m, backend.NewItem(backend.Category, true, nil, nil)
			return m, backend.NewItem(true, nil, nil)

		case "ctrl+e":
			// Edit the selected category
			if !m.list.IsEmpty() {
				categoryPath := []string{m.list.SelectedItem().FilterValue()}
				item := m.list.SelectedItem().(simplelist.Item)
				//TODO
				// return m, backend.NewItem(
				// 	backend.Category,
				// 	false,
				// 	categoryPath,
				// 	[]string{item.Title(), item.Description()},
				// )
				return m, backend.NewItem(
					false,
					categoryPath,
					[]string{item.Title(), item.Description()},
				)
			}

		//case "ctrl+d":
		// Delete the selected category
		//if !m.list.IsEmpty() {
		//TODO
		//return m, backend.DeleteItem(backend.Category, m.list.SelectedItem().FilterValue())
		//	return m, backend.DeleteItem(m.list.SelectedItem().FilterValue())
		//}

		default:
			// Check if we need to open a new category
			//TODO
			if item, ok := m.list.GetItem(msg.String()); ok {
				return m, tab.NewTab(item.FilterValue(), tab.Filter)
			}
		}
	}

	// Update the list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View the tab
func (m Model) View() string {
	// Check if the program is loaded, if not, return a loading message
	if !m.loaded {
		return "Loading..."
	}

	// Return the view
	return m.list.View()
}
