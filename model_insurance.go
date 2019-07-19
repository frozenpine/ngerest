package ngerest

// Insurance Fund Data
type Insurance struct {
	Currency string `json:"currency"`
	Timestamp Time `json:"timestamp"`
	WalletBalance float32 `json:"walletBalance,omitempty"`
}
