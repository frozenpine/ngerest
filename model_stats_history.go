package ngerest

import (
	"time"
)

// StatsHistory history states
type StatsHistory struct {
	Date       Time `json:"date"`
	RootSymbol string    `json:"rootSymbol"`
	Currency   string    `json:"currency,omitempty"`
	Volume     float32   `json:"volume,omitempty"`
	Turnover   float32   `json:"turnover,omitempty"`
}
