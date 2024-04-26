package history

import "time"

const (
	PositiveType = iota
	NegativeType
)

type History struct {
	TransactionID string    `json:"transaction_id"   db:"transaction_id"`
	WalletID      string    `json:"wallet_id"        db:"wallet_id"`
	UniqueID      string    `json:"unique_id"        db:"unique_id"`
	Type          int       `json:"type"             db:"type"`
	Amount        float64   `json:"amount"           db:"amount"`
	Notes         string    `json:"notes"            db:"notes"`
	Metadata      string    `json:"metadata"         db:"metadata"`
	CreatedAt     time.Time `json:"created_at"       db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"       db:"updated_at"`
}
