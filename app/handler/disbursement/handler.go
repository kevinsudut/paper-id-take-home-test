package disbursement

import (
	"context"

	"kevinsudut.com/paper-id-take-home-test/app/usecase/disbursement"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

type DisbursementUsecaseItf interface {
	DisbursementBalance(ctx context.Context, req disbursement.DisbursementBalanceRequest) (disbursement.DisbursementBalanceResponse, error)
}

type handler struct {
	logger       log.Logger
	disbursement DisbursementUsecaseItf
}

func InitHandler(logger log.Logger, disbursement DisbursementUsecaseItf) handler {
	return handler{
		logger:       logger,
		disbursement: disbursement,
	}
}
