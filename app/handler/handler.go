package handler

import (
	"net/http"

	"kevinsudut.com/paper-id-take-home-test/app/handler/disbursement"
	"kevinsudut.com/paper-id-take-home-test/app/usecase"
	httplib "kevinsudut.com/paper-id-take-home-test/pkg/lib/http"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/log"
)

type handler struct {
	handlers []httplib.Handler
}

func InitHandler(logger log.Logger, usecase usecase.Usecase) *handler {
	disbursement := disbursement.InitHandler(logger, usecase.Disbursement)

	return &handler{
		handlers: []httplib.Handler{
			{
				Path:   "/disbursement",
				Method: http.MethodPost,
				Func:   disbursement.DisbursementBalance,
			},
		},
	}
}

func (h handler) GetHandlers() []httplib.Handler {
	return h.handlers
}
