package domain

import (
	"kevinsudut.com/paper-id-take-home-test/app/domain/history"
	"kevinsudut.com/paper-id-take-home-test/app/domain/wallet"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
)

type Domain struct {
	Wallet  wallet.WalletDomainItf
	History history.HistoryDomainItf
}

func InitDomain(db database.DB) Domain {
	return Domain{
		Wallet:  wallet.InitDomain(db),
		History: history.InitDomain(db),
	}
}
