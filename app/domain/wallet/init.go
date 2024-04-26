package wallet

import (
	"context"
	"database/sql"

	historyresource "kevinsudut.com/paper-id-take-home-test/app/domain/history/database"
	resource "kevinsudut.com/paper-id-take-home-test/app/domain/wallet/database"
	historyentity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	entity "kevinsudut.com/paper-id-take-home-test/app/entity/wallet"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
)

type DBResourceItf interface {
	Begin() (*sql.Tx, error)
	GetWalletByWalletID(ctx context.Context, walletID string) (entity.Wallet, error)
	GrantWalletByWalletID(ctx context.Context, tx *sql.Tx, walletID string, amount float64) error
	DeductWalletByWalletID(ctx context.Context, tx *sql.Tx, walletID string, amount float64) error
}

type HistoryDBResourceItf interface {
	InsertHistory(ctx context.Context, tx *sql.Tx, data historyentity.History) error
}

type domain struct {
	db        DBResourceItf
	historyDb HistoryDBResourceItf
}

type WalletDomainItf interface {
	GetWalletByWalletID(ctx context.Context, walletID string) (entity.Wallet, error)
	DisbursementBalance(ctx context.Context, current DisbursementBalance, target DisbursementBalance) error
}

func InitDomain(db database.DB) domain {
	return domain{
		db:        resource.InitResource(db),
		historyDb: historyresource.InitResource(db),
	}
}
