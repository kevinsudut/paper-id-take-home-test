package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Path   string
	Method string
	Func   httprouter.Handle
}

func InitHttp(config Config, handlers []Handler) error {
	config.setDefaultConfig()

	router := httprouter.New()

	for _, handler := range handlers {
		switch handler.Method {
		case http.MethodGet:
			router.GET(handler.Path, handler.Func)
		case http.MethodPost:
			router.POST(handler.Path, handler.Func)
		case http.MethodPatch:
			router.PATCH(handler.Path, handler.Func)
		case http.MethodPut:
			router.PUT(handler.Path, handler.Func)
		}
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
}
