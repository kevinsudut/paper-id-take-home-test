package response

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"kevinsudut.com/paper-id-take-home-test/pkg/lib/errors"
)

type resultStatus struct {
	Code     string   `json:"code,omitempty"`
	Reason   string   `json:"reason,omitempty"`
	Messages []string `json:"messages,omitempty"`
	httpcode int      `json:"-"`
}

type result struct {
	ResultStatus resultStatus `json:"result_status"`
	Data         interface{}  `json:"data,omitempty"`
}

func WriteResponse(w http.ResponseWriter, data interface{}, err error) error {
	status := resultStatus{
		Code:     "200",
		Reason:   "Success",
		httpcode: http.StatusOK,
	}

	if err != nil {
		errs, ok := err.(*errors.Errs)
		if !ok || errs.Code == "" {
			errs = &errors.InternalServerError
			errs.Messages = []string{
				err.Error(),
			}
		}

		status.Code = errs.Code
		status.Reason = errs.Reason
		status.Messages = errs.Messages
		status.httpcode = errs.HttpCode
	}

	resp, err := jsoniter.Marshal(result{
		ResultStatus: status,
		Data:         data,
	})
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.httpcode)
	w.Write(resp)

	return nil
}
