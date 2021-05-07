package exception

import (
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/pkg/common/response"
	"net/http"
)

type RuntimeError struct {
	Code string
	Msg  string
	Data interface{}
	Err  error
}

func (re *RuntimeError) Error() string {
	return re.Code+","+re.Msg+","+re.Err.Error()
}

func (re *RuntimeError) Unwrap() error{
	return re.Err
}

func (re *RuntimeError) ToR() *response.R {
	return re.ToRWithStatusCode(http.StatusOK)
}

func (re *RuntimeError) ToRWithStatusCode(statusCode int) *response.R {
	return &response.R{
		StatusCode: statusCode,
		Code:       re.Code,
		Msg:        re.Msg,
		Data:       re.Data,
	}
}

