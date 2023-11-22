package data

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v56/github"
	"github.com/richardcase/herald/internal/utils"
)

func newGithubBackend() (NotificationBackend, error) {
	return &githhubBackend{}, nil
}

type githhubBackend struct {
	path string
}

func (g *githhubBackend) FetchNotifications(ctx context.Context, query string, limit int, pageInfo *PageInfo) (*NotificationsResponse, error) {
	token := os.Getenv("GITHUB_TOKEN")

	client := github.NewClient(nil).WithAuthToken(token)

	//TODO: handle other options like paging
	opts := &github.NotificationListOptions{
		All: true,
	}

	notifications, _, err := client.Activity.ListNotifications(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("getting notifications from github: %w", err)
	}

	mapped := fromGitHubNotifications(notifications)

	return &NotificationsResponse{
		Notifications: mapped,
		TotalCount:    len(mapped),
		PageInfo:      PageInfo{}, //TODO
	}, nil
}

func fromGitHubNotifications(gh []*github.Notification) []*NotificationData {
	ret := []*NotificationData{}

	for _, ghNotification := range gh {
		ret = append(ret, &NotificationData{
			ID:         utils.StringDeref(ghNotification.ID, ""),
			Reason:     Reason(utils.StringDeref(ghNotification.Reason, "")),
			Unread:     utils.BoolDeref(ghNotification.Unread, true),
			UpdatedAt:  ghNotification.UpdatedAt.Time,
			LastReadAt: &ghNotification.LastReadAt.Time,
			URL:        utils.StringDeref(ghNotification.URL, ""),
			Repository: Repository{
				Name:       utils.StringDeref(ghNotification.Repository.Name, ""),
				FullName:   utils.StringDeref(ghNotification.Repository.FullName, ""),
				IsArchived: utils.BoolDeref(ghNotification.Repository.Archived, false),
			},
			Subject: Subject{
				Title:            utils.StringDeref(ghNotification.Subject.Title, ""),
				URL:              utils.StringDeref(ghNotification.Subject.URL, ""),
				LatestCommentURL: utils.StringDeref(ghNotification.Subject.LatestCommentURL, ""),
				Type:             utils.StringDeref(ghNotification.Subject.Type, ""),
			},
		})
	}

	return ret
}
