package database

import (
	"context"
	"database/sql"
	"time"

	"kevinsudut.com/paper-id-take-home-test/app/entity/wallet"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (rsc resource) Begin() (*sql.Tx, error) {
	return rsc.db.Begin()
}

func (rsc resource) GetWalletByWalletID(ctx context.Context, walletID string) (resp wallet.Wallet, err error) {
	err = rsc.db.GetContext(ctx, &resp, queryGetWalletByWalletID, walletID)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	return resp, nil
}

func (rsc resource) GrantWalletByWalletID(ctx context.Context, tx *sql.Tx, walletID string, amount float64) (err error) {
	result, err := rsc.db.ExecContextTx(ctx, tx, queryGrantWalletByWalletID, amount, time.Now(), walletID)
	if err != nil {
		return errors.AddTrace(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.AddTrace(err)
	}

	if affected <= 0 {
		return errors.AddTrace(errors.FailedUpdateDB)
	}

	return nil
}

func (rsc resource) DeductWalletByWalletID(ctx context.Context, tx *sql.Tx, walletID string, amount float64) (err error) {
	result, err := rsc.db.ExecContextTx(ctx, tx, queryDeductWalletByWalletID, amount, time.Now(), walletID)
	if err != nil {
		return errors.AddTrace(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.AddTrace(err)
	}

	if affected <= 0 {
		return errors.AddTrace(errors.FailedUpdateDB)
	}

	return nil
}
