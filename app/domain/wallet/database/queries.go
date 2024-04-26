package database

const (
	queryGetWalletByWalletID = `
		SELECT
			wallet_id,
			user_id,
			amount,
			created_at,
			COALESCE(updated_at, created_at) as updated_at
		FROM
			wallets
		WHERE
			wallet_id = $1
	`

	queryGrantWalletByWalletID = `
		UPDATE wallets SET
			amount = amount + $1,
			updated_at = $2
		WHERE
			wallet_id = $3
	`

	queryDeductWalletByWalletID = `
		UPDATE wallets SET
			amount = amount - $1,
			updated_at = $2
		WHERE
			wallet_id = $3
		AND
			amount - $1 >= 0
	`
)
