package database

const (
	queryGetHistoryByUniqueID = `
		SELECT
			transaction_id,
			wallet_id,
			unique_id,
			type,
			amount,
			notes,
			COALESCE(metadata, '{}') as metadata,
			created_at,
			COALESCE(updated_at, created_at) as updated_at
		FROM
			histories
		WHERE
			unique_id = $1
	`

	queryInsertHistory = `
		INSERT INTO histories (
			transaction_id,
			wallet_id,
			unique_id,
			type,
			amount,
			notes,
			metadata,
			created_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		)
	`
)
