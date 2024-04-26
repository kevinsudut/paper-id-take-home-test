package disbursement

import (
	"context"
	"fmt"

	historyentity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (r DisbursementBalanceRequest) Validate(ctx context.Context, wallet WalletDomainItf, history HistoryDomainItf) error {
	missingErrs := errors.AddTrace(errors.MissingParameter)
	message := []string{}

	if r.WalletID == "" {
		message = append(message, "WallerID is required")
	}

	if r.UserID == "" {
		message = append(message, "UserID is required")
	}

	if r.UniqueID == "" {
		message = append(message, "UniqueID is required")
	}

	if r.TargetWalletID == "" {
		message = append(message, "TargetWalletID is required")
	}

	if r.Amount <= 0 {
		message = append(message, "Amount is required")
	}

	if len(message) > 0 {
		missingErrs.Messages = message
		return errors.AddTrace(missingErrs)
	}

	if r.WalletID == r.TargetWalletID {
		return errors.AddTrace(errors.CanNotSendMoneySameWallet)
	}

	userWallet, err := wallet.GetWalletByWalletID(ctx, r.WalletID)
	if err != nil && !errors.IsNotFound(err) {
		return errors.AddTrace(err)
	}

	if userWallet.UserID == "" {
		return errors.AddTrace(errors.WalletNotFound)
	}

	if userWallet.UserID != r.UserID {
		return errors.AddTrace(errors.WalletNotBelongToUser)
	}

	if userWallet.Amount < r.Amount {
		return errors.AddTrace(errors.BalanceIsNotEnough)
	}

	targetWallet, err := wallet.GetWalletByWalletID(ctx, r.TargetWalletID)
	if err != nil && !errors.IsNotFound(err) {
		return errors.AddTrace(err)
	}

	if targetWallet.UserID == "" {
		return errors.AddTrace(errors.TargetWalletNotFound)
	}

	historyResp, err := history.GetHistoryByUniqueID(ctx, fmt.Sprintf("%s-%d", r.UniqueID, historyentity.NegativeType))
	if err != nil && !errors.IsNotFound(err) {
		return errors.AddTrace(err)
	}

	if historyResp.TransactionID != "" {
		return errors.AddTrace(errors.UniqueIDAlreadyUsed)
	}

	return nil
}
