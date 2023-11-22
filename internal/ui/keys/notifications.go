package keys

import "github.com/charmbracelet/bubbles/key"

type NotificationsKeyMap struct {
	Read   key.Binding
	Unread key.Binding
}

var NotificationKeys = NotificationsKeyMap{
	Read: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "read"),
	),
	Unread: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "unread"),
	),
}

func NotificationsFullHelp() []key.Binding {
	return []key.Binding{
		NotificationKeys.Read,
		NotificationKeys.Unread,
	}
}
