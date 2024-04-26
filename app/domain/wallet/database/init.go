package database

import (
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
)

type resource struct {
	db database.DB
}

func InitResource(db database.DB) resource {
	return resource{
		db: db,
	}
}
