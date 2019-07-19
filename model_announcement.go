package ngerest

import (
	"time"
)

// Announcement public Announcements
type Announcement struct {
	ID      float32   `json:"id"`
	Link    string    `json:"link,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Date    time.Time `json:"date,omitempty"`
}
