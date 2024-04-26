package wallet

import (
	"context"

	entity "kevinsudut.com/paper-id-take-home-test/app/entity/wallet"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (d domain) GetWalletByWalletID(ctx context.Context, walletID string) (entity.Wallet, error) {
	resp, err := d.db.GetWalletByWalletID(ctx, walletID)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	return resp, nil
}

func (d domain) DisbursementBalance(ctx context.Context, current DisbursementBalance, target DisbursementBalance) error {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.AddTrace(err)
	}

	err = database.ExecuteTx(tx, func() error {
		// deduct current wallet
		err = d.db.DeductWalletByWalletID(ctx, tx, current.Wallet.WalletID, current.Wallet.Amount)
		if err != nil {
			return errors.AddTrace(err)
		}

		err = d.historyDb.InsertHistory(ctx, tx, current.History)
		if err != nil {
			return errors.AddTrace(err)
		}

		// grant target wallet
		err = d.db.GrantWalletByWalletID(ctx, tx, target.Wallet.WalletID, target.Wallet.Amount)
		if err != nil {
			return errors.AddTrace(err)
		}

		err = d.historyDb.InsertHistory(ctx, tx, target.History)
		if err != nil {
			return errors.AddTrace(err)
		}

		return nil
	})
	if err != nil {
		return errors.AddTrace(err)
	}

	return nil
}
