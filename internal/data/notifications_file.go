package data

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func newFileBackend(path string) (NotificationBackend, error) {
	return &fileBackend{
		path: path,
	}, nil
}

type fileBackend struct {
	path string
}

func (fb *fileBackend) FetchNotifications(ctx context.Context, query string, limit int, pageInfo *PageInfo) (*NotificationsResponse, error) {

	data, err := os.ReadFile(fb.path)
	if err != nil {
		return nil, fmt.Errorf("getting notifications from file %s: %w", fb.path, err)
	}

	items := []*NotificationData{}
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("marshalling data: %w", err)
	}

	return &NotificationsResponse{
		Notifications: items,
		TotalCount:    len(items),
		PageInfo:      PageInfo{}, //TODO
	}, nil
}
