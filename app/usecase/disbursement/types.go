package disbursement

type DisbursementBalanceRequest struct {
	UserID         string  `json:"user_id"`
	UniqueID       string  `json:"unique_id"`
	WalletID       string  `json:"wallet_id"`
	Amount         float64 `json:"amount"`
	Notes          string  `json:"notes"`
	TargetWalletID string  `json:"target_wallet_id"`
}

type DisbursementBalanceResponse struct {
	TransactionID string `json:"transaction_id"`
}
