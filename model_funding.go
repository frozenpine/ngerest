package ngerest

import (
	"time"
)

// Funding Swap Funding History
type Funding struct {
	Timestamp        Time `json:"timestamp"`
	Symbol           string    `json:"symbol"`
	FundingInterval  Time `json:"fundingInterval,omitempty"`
	FundingRate      float64   `json:"fundingRate,omitempty"`
	FundingRateDaily float64   `json:"fundingRateDaily,omitempty"`
}
