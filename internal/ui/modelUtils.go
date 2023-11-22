package ui

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/richardcase/herald/internal/config"
	"github.com/richardcase/herald/internal/data"
	"github.com/richardcase/herald/internal/ui/common"
	"github.com/richardcase/herald/internal/ui/components/section"
	"github.com/richardcase/herald/internal/ui/constants"
	"github.com/richardcase/herald/internal/ui/context"
	"github.com/richardcase/herald/internal/ui/markdown"
)

func (m *Model) getCurrSection() section.Section {
	sections := m.getCurrentViewSections()
	if len(sections) == 0 || m.currSectionId >= len(sections) {
		return nil
	}
	return sections[m.currSectionId]
}

func (m *Model) getCurrRowData() data.RowData {
	section := m.getCurrSection()
	if section == nil {
		return nil
	}
	return section.GetCurrRow()
}

func (m *Model) getSectionAt(id int) section.Section {
	sections := m.getCurrentViewSections()
	if len(sections) <= id {
		return nil
	}
	return sections[id]
}

func (m *Model) getPrevSectionId() int {
	sectionsConfigs := m.ctx.GetViewSectionsConfig()
	m.currSectionId = (m.currSectionId - 1) % len(sectionsConfigs)
	if m.currSectionId < 0 {
		m.currSectionId += len(sectionsConfigs)
	}

	return m.currSectionId
}

func (m *Model) getNextSectionId() int {
	return (m.currSectionId + 1) % len(m.ctx.GetViewSectionsConfig())
}

type NotificationsCommandTemplateInput struct {
	RepoName string
	RepoPath string
}

func (m *Model) executeKeybinding(key string) tea.Cmd {
	currRowData := m.getCurrRowData()

	switch m.ctx.View {
	case config.NotificationsView:
		for _, keybinding := range m.ctx.Config.Keybindings.Notifications {
			if keybinding.Key != key {
				continue
			}

			switch data := currRowData.(type) {
			case *data.NotificationData:
				return m.runCustomNotificationsCommand(keybinding.Command, data)
			}
		}
	default:
		// Not a valid case - ignore it
	}

	return nil
}

func (m *Model) runCustomNotificationsCommand(commandTemplate string, notificationData *data.NotificationData) tea.Cmd {
	repoName := notificationData.GetRepoName()
	repoPath, ok := common.GetRepoLocalPath(repoName, m.ctx.Config.RepoPaths)

	if !ok {
		return func() tea.Msg {
			return constants.ErrMsg{Err: fmt.Errorf("Failed to find local path for repo %s", repoName)}
		}
	}

	input := NotificationsCommandTemplateInput{
		RepoName: repoName,
		RepoPath: repoPath,
	}

	cmd, err := template.New("keybinding_command").Parse(commandTemplate)
	if err != nil {
		log.Fatal("Failed parse keybinding template", err)
	}

	var buff bytes.Buffer
	err = cmd.Execute(&buff, input)
	if err != nil {
		log.Fatal("Failed executing keybinding command", err)
	}
	return m.executeCustomCommand(buff.String())
}

func (m *Model) executeCustomCommand(cmd string) tea.Cmd {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}
	c := exec.Command(shell, "-c", cmd)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			mdRenderer := markdown.GetMarkdownRenderer(m.ctx.ScreenWidth)
			md, mdErr := mdRenderer.Render(fmt.Sprintf("While running: `%s`", cmd))
			if mdErr != nil {
				return constants.ErrMsg{Err: mdErr}
			}
			return constants.ErrMsg{Err: fmt.Errorf(
				lipgloss.JoinVertical(lipgloss.Left,
					fmt.Sprintf("Whoops, got an error: %s", err),
					md,
				),
			)}
		}
		return nil
	})
}

func (m *Model) notify(text string) tea.Cmd {
	id := fmt.Sprint(time.Now().Unix())
	m.tasks[id] = context.Task{
		Id:           id,
		FinishedText: text,
		State:        context.TaskFinished,
	}
	return func() tea.Msg {
		return constants.TaskFinishedMsg{
			SectionId:   m.getCurrSection().GetId(),
			SectionType: m.getCurrSection().GetType(),
			TaskId:      id,
			Err:         nil,
		}
	}
}
