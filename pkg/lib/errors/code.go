package errors

import (
	"net/http"
)

var (
	BadRequest = Errs{
		Code:     "400",
		Reason:   "Bad Request",
		HttpCode: http.StatusBadRequest,
	}

	MissingParameter = Errs{
		Code:     "400",
		Reason:   "Missing Required Fields",
		HttpCode: http.StatusBadRequest,
	}

	CanNotSendMoneySameWallet = Errs{
		Code:     "403",
		Reason:   "Can not send money to same wallet",
		HttpCode: http.StatusForbidden,
	}

	WalletNotFound = Errs{
		Code:     "403",
		Reason:   "User wallet not found",
		HttpCode: http.StatusForbidden,
	}

	WalletNotBelongToUser = Errs{
		Code:     "403",
		Reason:   "Wallet not belong to current user",
		HttpCode: http.StatusForbidden,
	}

	BalanceIsNotEnough = Errs{
		Code:     "403",
		Reason:   "Balance is not enough",
		HttpCode: http.StatusForbidden,
	}

	TargetWalletNotFound = Errs{
		Code:     "403",
		Reason:   "Target user wallet not found",
		HttpCode: http.StatusForbidden,
	}

	UniqueIDAlreadyUsed = Errs{
		Code:     "409",
		Reason:   "Unique id already used",
		HttpCode: http.StatusConflict,
	}

	InternalServerError = Errs{
		Code:     "500",
		Reason:   "Internal Server Error",
		HttpCode: http.StatusInternalServerError,
	}

	FailedInsertDB = Errs{
		Code:     "500",
		Reason:   "Failed Insert DB",
		HttpCode: http.StatusInternalServerError,
	}

	FailedUpdateDB = Errs{
		Code:     "500",
		Reason:   "Failed Update DB",
		HttpCode: http.StatusInternalServerError,
	}
)
