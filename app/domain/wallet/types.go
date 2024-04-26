package wallet

import (
	historyentity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	entity "kevinsudut.com/paper-id-take-home-test/app/entity/wallet"
)

type DisbursementBalance struct {
	Wallet  entity.Wallet
	History historyentity.History
}
