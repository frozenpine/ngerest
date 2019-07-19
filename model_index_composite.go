package ngerest

import (
	"time"
)

// IndexComposite index composite
type IndexComposite struct {
	Timestamp   Time `json:"timestamp"`
	Symbol      string    `json:"symbol,omitempty"`
	IndexSymbol string    `json:"indexSymbol,omitempty"`
	Reference   string    `json:"reference,omitempty"`
	LastPrice   float64   `json:"lastPrice,omitempty"`
	Weight      float64   `json:"weight,omitempty"`
	Logged      Time `json:"logged,omitempty"`
}
