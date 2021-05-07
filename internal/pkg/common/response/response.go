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
var InternalServerError = New(http.StatusInternalServerError,"SYS_500","服务器内部错误",nil)
func New(statusCode int, code, msg string, data interface{}) *R {
	return &R{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
		Data:       data,
	}
}

func (r *R) Unwrap()(int,*R){
	return r.StatusCode,r
}

func BuildSuccessRWithNoData()*R{
	return New(http.StatusOK,"SYS_200","请求成功",nil)
}
func Json(ctx *gin.Context, r *R) {
	ctx.JSON(r.StatusCode, r)
}
func AbortWithError(ctx *gin.Context,err error){
	ctx.Abort()
	_ = ctx.Error(err)
}


func SuccessWithNoData(ctx *gin.Context){
	ctx.JSON(BuildSuccessRWithNoData().Unwrap())
}

func Success(ctx *gin.Context,data interface{}){
	r := BuildSuccessRWithNoData()
	r.Data = data
	ctx.JSON(r.Unwrap())
}
