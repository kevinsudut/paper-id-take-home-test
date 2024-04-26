package disbursement

import (
	"context"

	"kevinsudut.com/paper-id-take-home-test/app/domain/wallet"
	historyentity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	walletentity "kevinsudut.com/paper-id-take-home-test/app/entity/wallet"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

type WalletDomainItf interface {
	GetWalletByWalletID(ctx context.Context, walletID string) (walletentity.Wallet, error)
	DisbursementBalance(ctx context.Context, current wallet.DisbursementBalance, target wallet.DisbursementBalance) error
}

type HistoryDomainItf interface {
	GetHistoryByUniqueID(ctx context.Context, uniqueID string) (historyentity.History, error)
}

type usecase struct {
	logger  log.Logger
	wallet  WalletDomainItf
	history HistoryDomainItf
}

type DisbursementUsecaseItf interface {
	DisbursementBalance(ctx context.Context, req DisbursementBalanceRequest) (DisbursementBalanceResponse, error)
}

func InitUsecase(logger log.Logger, wallet WalletDomainItf, history HistoryDomainItf) usecase {
	return usecase{
		logger:  logger,
		wallet:  wallet,
		history: history,
	}
}
