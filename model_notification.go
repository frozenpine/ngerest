package ngerest

import (
	"time"
)

// Notification Account Notifications
type Notification struct {
	ID                float32   `json:"id,omitempty"`
	Date              Time `json:"date"`
	Title             string    `json:"title"`
	Body              string    `json:"body"`
	TTL               float32   `json:"ttl"`
	Type              string    `json:"type,omitempty"`
	Closable          bool      `json:"closable,omitempty"`
	Persist           bool      `json:"persist,omitempty"`
	WaitForVisibility bool      `json:"waitForVisibility,omitempty"`
	Sound             string    `json:"sound,omitempty"`
}
