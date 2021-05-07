package middleware

import (
	"errors"
	"fmt"
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/exception"
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/pkg/common/response"
	"github.com/gin-gonic/gin"
	"log"
)

func GlobalErrorHandler(ctx *gin.Context) {
	defer func() {
		size := len(ctx.Errors)
		if size > 0 {
			ginErr := ctx.Errors[size-1]
			var err *exception.RuntimeError
			if errors.As(ginErr.Err, &err) {
				//用户系统自定义的错误
				response.Json(ctx, err.ToR())
			} else {
				//其他错误
				//TODO 根据业务增加其他错误的处理
				log.Printf("\nRequest:\n%s\nError:\n%+v\n", formatRequest(ctx), ginErr.Err)
				response.Json(ctx, response.InternalServerError)
			}
			//避免gin重复输出这个错误
			ctx.Errors = ctx.Errors[:size-1]
		}
	}()
	ctx.Next()
}

func formatRequest(ctx *gin.Context) string {
	r := ctx.Request
	result := r.Method + " " + r.RequestURI + " " + r.Proto + "\n"
	result += "Header:\n"
	for k, v := range r.Header {
		result += fmt.Sprintf("%s: %s\n", k, v)
	}
	result += "Body:\n"
	body, _ := ctx.Get("RequestBody")
	if b, ok := body.(string); ok {
		result += b
	}
	return result
}