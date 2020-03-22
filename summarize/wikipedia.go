package summarize

import (
	"fmt"

	"framagit.org/andinus/indus/fetch"
	"framagit.org/andinus/indus/notification"
)

// Wikipedia returns struct notification.Notif with notification info
// from the wikipedia summary.
func Wikipedia(w fetch.Wiki) (notification.Notif, error) {
	n := notification.Notif{}
	var err error

	// Continue only if the page's type is standard. TODO: Work
	// with other page types to get summary.
	switch w.Type {
	case "standard":
		n.Title = fmt.Sprintf("%s", w.Title)
		n.Message = w.Extract
	default:
		err = fmt.Errorf("Summarizing wikipedia response failed")
	}

	return n, err
}
