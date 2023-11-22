package pr

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/richardcase/herald/internal/data"
	"github.com/richardcase/herald/internal/ui/components"
	"github.com/richardcase/herald/internal/ui/components/table"
	"github.com/richardcase/herald/internal/ui/context"
	"github.com/richardcase/herald/internal/utils"
)

type Notification struct {
	Ctx  *context.ProgramContext
	Data data.NotificationData
}

type sectionNotificationFetchedMsg struct {
	SectionId     int
	Notifications []Notification
}

func (n *Notification) getTextStyle() lipgloss.Style {
	return components.GetNotificationTextStyle(n.Ctx, n.Data.Unread)
}

//func (n *Notification) renderReviewStatus() string {
//reviewCellStyle := n.getTextStyle()
// if pr.Data.ReviewDecision == "APPROVED" {
// 	if pr.Data.State == "OPEN" {
// 		reviewCellStyle = reviewCellStyle.Foreground(
// 			pr.Ctx.Theme.SuccessText,
// 		)
// 	}
// 	return reviewCellStyle.Render("󰄬")
// }

// if pr.Data.ReviewDecision == "CHANGES_REQUESTED" {
// 	if pr.Data.State == "OPEN" {
// 		reviewCellStyle = reviewCellStyle.Foreground(
// 			pr.Ctx.Theme.WarningText,
// 		)
// 	}
// 	return reviewCellStyle.Render("󰌑")
// }

//	return reviewCellStyle.Render(n.Ctx.Styles.Common.WaitingGlyph)
//}

func (n *Notification) renderReason() string {
	mergeCellStyle := lipgloss.NewStyle()
	switch n.Data.Reason {
	//case data.Author:
	// 	if pr.Data.IsDraft {
	// 		return mergeCellStyle.Foreground(n.Ctx.Theme.FaintText).Render("")
	// 	} else {
	// 		return mergeCellStyle.Foreground(n.Ctx.Styles.Colors.OpenPR).Render("")
	// 	}
	// case "CLOSED":
	// 	return mergeCellStyle.Foreground(n.Ctx.Styles.Colors.ClosedPR).
	// 		Render("")
	// case "MERGED":
	// 	return mergeCellStyle.Foreground(n.Ctx.Styles.Colors.MergedPR).
	// 		Render("")
	default:
		return mergeCellStyle.Foreground(n.Ctx.Theme.FaintText).Render("-")
	}
}

// func (n *Notification) GetStatusChecksRollup() string {
// 	if pr.Data.Mergeable == "CONFLICTING" {
// 		return "FAILURE"
// 	}

// 	accStatus := "SUCCESS"
// 	commits := pr.Data.Commits.Nodes
// 	if len(commits) == 0 {
// 		return "PENDING"
// 	}

// 	mostRecentCommit := commits[0].Commit
// 	for _, statusCheck := range mostRecentCommit.StatusCheckRollup.Contexts.Nodes {
// 		var conclusion string
// 		if statusCheck.Typename == "CheckRun" {
// 			conclusion = string(statusCheck.CheckRun.Conclusion)
// 			status := string(statusCheck.CheckRun.Status)
// 			if isStatusWaiting(status) {
// 				accStatus = "PENDING"
// 			}
// 		} else if statusCheck.Typename == "StatusContext" {
// 			conclusion = string(statusCheck.StatusContext.State)
// 			if isStatusWaiting(conclusion) {
// 				accStatus = "PENDING"
// 			}
// 		}

// 		if isConclusionAFailure(conclusion) {
// 			accStatus = "FAILURE"
// 			break
// 		}
// 	}

// 	return accStatus
// }

// func (pr *PullRequest) renderCiStatus() string {

// 	accStatus := pr.GetStatusChecksRollup()
// 	ciCellStyle := pr.getTextStyle()
// 	if accStatus == "SUCCESS" {
// 		if pr.Data.State == "OPEN" {
// 			ciCellStyle = ciCellStyle.Foreground(pr.Ctx.Theme.SuccessText)
// 		}
// 		return ciCellStyle.Render(constants.SuccessIcon)
// 	}

// 	if accStatus == "PENDING" {
// 		return ciCellStyle.Render(pr.Ctx.Styles.Common.WaitingGlyph)
// 	}

// 	if pr.Data.State == "OPEN" {
// 		ciCellStyle = ciCellStyle.Foreground(pr.Ctx.Theme.WarningText)
// 	}
// 	return ciCellStyle.Render(constants.FailureIcon)
// }

// func (n *Notification) renderLines(isSelected bool) string {
// 	deletions := 0
// 	if pr.Data.Deletions > 0 {
// 		deletions = pr.Data.Deletions
// 	}

// 	var additionsFg, deletionsFg lipgloss.AdaptiveColor
// 	state := pr.Data.State
// 	if state != "OPEN" {
// 		additionsFg = pr.Ctx.Theme.FaintText
// 		deletionsFg = pr.Ctx.Theme.FaintText
// 	} else {
// 		additionsFg = pr.Ctx.Theme.SuccessText
// 		deletionsFg = pr.Ctx.Theme.WarningText
// 	}

// 	baseStyle := lipgloss.NewStyle()
// 	if isSelected {
// 		baseStyle = baseStyle.Background(pr.Ctx.Theme.SelectedBackground)
// 	}

// 	additionsText := baseStyle.Copy().
// 		Foreground(additionsFg).
// 		Render(fmt.Sprintf("+%s", components.FormatNumber(pr.Data.Additions)))
// 	deletionsText := baseStyle.Copy().
// 		Foreground(deletionsFg).
// 		Render(fmt.Sprintf("-%s", components.FormatNumber(deletions)))

// 	return pr.getTextStyle().Render(
// 		keepSameSpacesOnAddDeletions(
// 			lipgloss.JoinHorizontal(
// 				lipgloss.Left,
// 				additionsText,
// 				baseStyle.Render(" "),
// 				deletionsText,
// 			)),
// 	)
// }

func keepSameSpacesOnAddDeletions(str string) string {
	strAsList := strings.Split(str, " ")
	return fmt.Sprintf(
		"%7s",
		strAsList[0],
	) + " " + fmt.Sprintf(
		"%7s",
		strAsList[1],
	)
}

func (n *Notification) renderTitle() string {
	return components.RenderNotificationTitle(
		n.Ctx,
		n.Data.Unread,
		n.Data.Subject.Title,
	)
}

// func (pr *PullRequest) renderAuthor() string {
// 	return pr.getTextStyle().Render(pr.Data.Author.Login)
// }

// func (pr *PullRequest) renderAssignees() string {
// 	assignees := make([]string, 0, len(pr.Data.Assignees.Nodes))
// 	for _, assignee := range pr.Data.Assignees.Nodes {
// 		assignees = append(assignees, assignee.Login)
// 	}
// 	return pr.getTextStyle().Render(strings.Join(assignees, ","))
// }

func (n *Notification) renderRepoName() string {
	repoName := n.Data.Repository.FullName
	return n.getTextStyle().Render(repoName)
}

func (n *Notification) renderUpdateAt() string {
	return n.getTextStyle().
		Render(utils.TimeElapsed(n.Data.UpdatedAt))
}

func (n *Notification) RenderReason() string {
	// switch n.Data.Reason {
	// case "OPEN":
	// 	if pr.Data.IsDraft {
	// 		return " Draft"
	// 	} else {
	// 		return " Open"
	// 	}
	// case "CLOSED":
	// 	return "󰗨 Closed"
	// case "MERGED":
	// 	return " Merged"
	// default:
	// 	return ""
	// }

	//TODO: chnage based on the reason
	return n.getTextStyle().Render(string(n.Data.Reason))
}

func (n *Notification) ToTableRow(isSelected bool) table.Row {
	return table.Row{
		n.renderUpdateAt(),
		n.renderReason(),
		n.renderRepoName(),
		n.renderTitle(),
	}
}

// func isConclusionAFailure(conclusion string) bool {
// 	return conclusion == "FAILURE" || conclusion == "TIMED_OUT" ||
// 		conclusion == "STARTUP_FAILURE"
// }

// func isStatusWaiting(status string) bool {
// 	return status == "PENDING" ||
// 		status == "QUEUED" ||
// 		status == "IN_PROGRESS" ||
// 		status == "WAITING"
// }
