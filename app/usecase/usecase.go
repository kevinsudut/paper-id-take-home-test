package usecase

import (
	"kevinsudut.com/paper-id-take-home-test/app/domain"
	"kevinsudut.com/paper-id-take-home-test/app/usecase/disbursement"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

type Usecase struct {
	Disbursement disbursement.DisbursementUsecaseItf
}

func InitUsecase(logger log.Logger, domain domain.Domain) Usecase {
	return Usecase{
		Disbursement: disbursement.InitUsecase(logger, domain.Wallet, domain.History),
	}
}
