package database

import (
	"context"
	"database/sql"
	"time"

	"kevinsudut.com/paper-id-take-home-test/app/entity/history"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (rsc resource) GetHistoryByUniqueID(ctx context.Context, uniqueID string) (resp history.History, err error) {
	err = rsc.db.GetContext(ctx, &resp, queryGetHistoryByUniqueID, uniqueID)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	return resp, nil
}

func (rsc resource) InsertHistory(ctx context.Context, tx *sql.Tx, data history.History) error {
	if data.Metadata == "" {
		data.Metadata = "{}"
	}

	result, err := rsc.db.ExecContextTx(ctx, tx, queryInsertHistory,
		data.TransactionID,
		data.WalletID,
		data.UniqueID,
		data.Type,
		data.Amount,
		data.Notes,
		data.Metadata,
		time.Now())
	if err != nil {
		return errors.AddTrace(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.AddTrace(err)
	}

	if affected <= 0 {
		return errors.AddTrace(errors.FailedInsertDB)
	}

	return nil
}
