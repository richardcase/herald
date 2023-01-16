package backend

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// FetchSuccessMessage is a message that is sent when the fetching of the
// notifications was successful
type FetchSuccessMessage struct {
	Items []list.Item
}

// FetchErrorMessage is a message that is sent when the fetching of the
// notifications failed
type FetchErrorMessage struct {
	Description string
	Err         error
}

// NewItemMessage is a message to tell the main model that a new item
// needs to be added to the list
type NewItemMessage struct {
	New       bool
	Fields    []string
	ItemPath  []string
	OldFields []string
}

// NewItem is a function to tell the main model that a new item
// needs to be added to the list
func NewItem(newItem bool, itemPath []string, oldFields []string) tea.Cmd {
	return func() tea.Msg {
		var textFields []string
		textFields = []string{"Name", "URL"}

		return NewItemMessage{
			New:       newItem,
			Fields:    textFields,
			ItemPath:  itemPath,
			OldFields: oldFields,
		}
	}
}

/*
// DeleteItemMessage is a message to tell the main model that a new item
// needs to be removed from the list
type DeleteItemMessage struct {
	Key string
}

// DeleteItem is a function to tell the main model that a new item
// needs to be removed from the list
func DeleteItem(key string) tea.Cmd {
	return func() tea.Msg {
		return DeleteItemMessage{
			Key: key,
		}
	}
}*/
