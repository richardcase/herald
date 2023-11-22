package data

import (
	"context"
	"time"
)

type NotificationsResponse struct {
	Notifications []*NotificationData
	TotalCount    int
	PageInfo      PageInfo
}

type NotificationData struct {
	ID         string     `json:"id"`
	Repository Repository `json:"repository"`
	Subject    Subject    `json:"subject"`
	Reason     Reason     `json:"reason"`
	Unread     bool       `json:"unread"`
	UpdatedAt  time.Time  `json:"updated_at"`
	LastReadAt *time.Time `json:"last_read_at"`
	URL        string     `json:"url,omitempty"`
}

func (n *NotificationData) GetRepoName() string {
	return n.Repository.FullName
}

func (n *NotificationData) GetID() string {
	return n.ID
}

func (n *NotificationData) GetUrl() string {
	return n.URL
}

func (n *NotificationData) GetUpdatedAt() time.Time {
	return n.UpdatedAt
}

func (n *NotificationData) GetReason() string {
	return string(n.Reason)
}

// From: https://docs.github.com/en/free-pro-team@latest/rest/activity/notifications?apiVersion=2022-11-28#about-notification-reasons
type Reason string

var (
	Assign          Reason = "assign"
	Author          Reason = "author"
	Comment         Reason = "comment"
	CIActivity      Reason = "ci_activity"
	Invitation      Reason = "invitation"
	Manual          Reason = "manual"
	Mention         Reason = "mention"
	ReviewRequested Reason = "review_requested"
	SecurityAlert   Reason = "security_alert"
	StateChange     Reason = "state_change"
	Subscribed      Reason = "subscribed"
	TeamMention     Reason = "team_mention"
)

type Repository struct {
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	IsArchived bool   `json:"archived"`
}

type Subject struct {
	Title            string `json:"title"`
	URL              string `json:"url,omitempty"`
	LatestCommentURL string `json:"latest_comment_url"`
	Type             string `json:"type"`
}

type NotificationBackend interface {
	FetchNotifications(ctx context.Context, query string, limit int, pageInfo *PageInfo) (*NotificationsResponse, error)
}

type Type int

const (
	File Type = iota
	Github
)

type PageInfo struct {
	HasNextPage bool
	StartCursor string
	EndCursor   string
}

func New(backedType Type) (NotificationBackend, error) {
	switch backedType {
	case File:
		return newFileBackend("_test/test.json")
	default:
		return newGithubBackend()
	}
}
