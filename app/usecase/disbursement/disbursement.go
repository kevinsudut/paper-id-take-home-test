package disbursement

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"kevinsudut.com/paper-id-take-home-test/app/domain/wallet"
	historyentity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	walletentity "kevinsudut.com/paper-id-take-home-test/app/entity/wallet"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (uc usecase) DisbursementBalance(ctx context.Context, req DisbursementBalanceRequest) (resp DisbursementBalanceResponse, err error) {
	err = req.Validate(ctx, uc.wallet, uc.history)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	transactionID := uuid.NewString()

	jsonReq, err := jsoniter.MarshalToString(req)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	err = uc.wallet.DisbursementBalance(ctx, wallet.DisbursementBalance{
		Wallet: walletentity.Wallet{
			WalletID: req.WalletID,
			Amount:   req.Amount,
		},
		History: historyentity.History{
			TransactionID: transactionID,
			WalletID:      req.WalletID,
			UniqueID:      fmt.Sprintf("%s-%d", req.UniqueID, historyentity.NegativeType),
			Type:          historyentity.NegativeType,
			Amount:        req.Amount,
			Notes:         req.Notes,
			Metadata:      jsonReq,
		},
	}, wallet.DisbursementBalance{
		Wallet: walletentity.Wallet{
			WalletID: req.TargetWalletID,
			Amount:   req.Amount,
		},
		History: historyentity.History{
			TransactionID: uuid.NewString(),
			WalletID:      req.TargetWalletID,
			UniqueID:      fmt.Sprintf("%s-%d", req.UniqueID, historyentity.PositiveType),
			Type:          historyentity.PositiveType,
			Amount:        req.Amount,
			Notes:         fmt.Sprintf("Receive money from %s", req.WalletID),
			Metadata:      jsonReq,
		},
	})
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	return DisbursementBalanceResponse{
		TransactionID: transactionID,
	}, nil
}
