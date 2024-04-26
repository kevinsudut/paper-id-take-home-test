package app

import (
	"kevinsudut.com/paper-id-take-home-test/app/domain"
	"kevinsudut.com/paper-id-take-home-test/app/handler"
	"kevinsudut.com/paper-id-take-home-test/app/usecase"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/database"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/http"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

func InitApp(config http.Config, logger log.Logger, db database.DB) error {
	usecase := usecase.InitUsecase(logger, domain.InitDomain(db))

	return http.InitHttp(config, handler.InitHandler(logger, usecase).GetHandlers())
}
