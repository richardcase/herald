package backend

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v49/github"
	"github.com/richardcase/herald/internal/model/simplelist"
)

func newFileBackend(path string) (Backend, error) {
	return &fileBackend{
		path: path,
	}, nil
}

type fileBackend struct {
	path string
}

func (f *fileBackend) FetchNotifications(repoFilters []string) tea.Cmd {
	return func() tea.Msg {
		items := []*github.Notification{}

		data, err := os.ReadFile(f.path)
		if err != nil {
			return FetchErrorMessage{
				Description: fmt.Sprintf("reading notifications from files %s", f.path),
				Err:         err,
			}
		}

		if err := json.Unmarshal(data, &items); err != nil {
			return FetchErrorMessage{
				Description: "unmarshalling notifications from file",
				Err:         err,
			}
		}

		//TODO: filter the repos

		var result []list.Item
		for _, item := range items {
			result = append(result, simplelist.NewItem(
				*item.Subject.Title,
				*item.Repository.FullName,
				*item.Subject.URL,
			))
		}

		return FetchSuccessMessage{Items: result}
	}
}

func (f *fileBackend) Close() error {
	return nil
}
