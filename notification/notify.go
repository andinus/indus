package notification

import (
	"os/exec"
)

// Notif struct holds information about the notification. Other
// parameters like urgency & timeout could be added when required.
type Notif struct {
	Title   string
	Message string
}

// Notify sends a desktop notification to the user using libnotify. It
// handles information in the form of Notif struct. It returns an
// error (if exists).
func (n Notif) Notify() error {
	err := exec.Command("notify-send", n.Title, n.Message).Run()
	return err
}
