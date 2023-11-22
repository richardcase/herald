package data

import "time"

type RowData interface {
	GetRepoName() string
	GetID() string
	GetUrl() string
	GetUpdatedAt() time.Time
	GetReason() string
}

func IsStatusWaiting(status string) bool {
	return status == "PENDING" ||
		status == "QUEUED" ||
		status == "IN_PROGRESS" ||
		status == "WAITING"
}

func IsConclusionAFailure(conclusion string) bool {
	return conclusion == "FAILURE" || conclusion == "TIMED_OUT" || conclusion == "STARTUP_FAILURE"
}
