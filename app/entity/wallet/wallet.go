package wallet

import "time"

type Wallet struct {
	WalletID  string    `json:"wallet_id"   db:"wallet_id"`
	UserID    string    `json:"user_id"     db:"user_id"`
	Amount    float64   `json:"amount"      db:"amount"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
	UpdatedAt time.Time `json:"updated_at"  db:"updated_at"`
}
