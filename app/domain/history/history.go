package history

import (
	"context"

	entity "kevinsudut.com/paper-id-take-home-test/app/entity/history"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (d domain) GetHistoryByUniqueID(ctx context.Context, uniqueID string) (entity.History, error) {
	resp, err := d.db.GetHistoryByUniqueID(ctx, uniqueID)
	if err != nil {
		return resp, errors.AddTrace(err)
	}

	return resp, nil
}
