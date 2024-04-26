package disbursement

import (
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"kevinsudut.com/paper-id-take-home-test/app/usecase/disbursement"
	"kevinsudut.com/paper-id-take-home-test/pkg/helper/response"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

func (h handler) DisbursementBalance(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		err  error
		req  disbursement.DisbursementBalanceRequest
		resp disbursement.DisbursementBalanceResponse
	)

	defer func() {
		if err != nil {
			h.logger.ErrorTrace(err, fmt.Sprintf("[Delivery][DisbursementBalance][Request:%+v]", req))
		}

		response.WriteResponse(w, resp, err)
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errs := errors.AddTrace(errors.BadRequest)
		errs.Messages = []string{err.Error()}
		return
	}
	defer r.Body.Close()

	err = jsoniter.Unmarshal(body, &req)
	if err != nil {
		errs := errors.AddTrace(errors.BadRequest)
		errs.Messages = []string{err.Error()}
		return
	}

	resp, err = h.disbursement.DisbursementBalance(r.Context(), req)
	if err != nil {
		return
	}
}
