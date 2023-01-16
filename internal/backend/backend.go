package backend

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Backend interface {
	FetchNotifications(repoFilters []string) tea.Cmd
	Close() error
}

type Type int

const (
	File Type = iota
	Web
)

func New(backedType Type) (Backend, error) {
	switch backedType {
	case File:
		return newFileBackend("_test/test.json")
	default:
		return nil, fmt.Errorf("type not supported")
	}
}
