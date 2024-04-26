package history

import (
	"context"

	resource "kevinsudut.com/paper-id-take-home-test/app/domain/history/database"
	entity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
)

type dbResourceItf interface {
	GetHistoryByUniqueID(ctx context.Context, uniqueID string) (entity.History, error)
}

type domain struct {
	db dbResourceItf
}

type HistoryDomainItf interface {
	GetHistoryByUniqueID(ctx context.Context, uniqueID string) (entity.History, error)
}

func InitDomain(db database.DB) domain {
	return domain{
		db: resource.InitResource(db),
	}
}
