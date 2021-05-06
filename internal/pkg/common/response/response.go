package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}


func New(statusCode int, code, msg string, data interface{}) *R {
	return &R{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
		Data:       data,
	}
}

var (
	Success                     = New(http.StatusOK, "SYS_200", "请求成功", nil)
	InternalServerError         = New(http.StatusInternalServerError, "SYS_500", "服务器内部错误", nil)
)

func Json(ctx *gin.Context, r *R) {
	ctx.JSON(r.StatusCode, r)
}

