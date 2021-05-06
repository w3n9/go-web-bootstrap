package exception

import (
	"errors"
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/pkg/common/response"
	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context) {
	defer func() {
		size := len(ctx.Errors)
		if size > 0 {
			if gin.Mode() == gin.ReleaseMode {
				response.Json(ctx, response.InternalServerError)
			}else{
				ginErr:=ctx.Errors[size-1]
				var err *RuntimeError
				if errors.As(ginErr.Err,&err){
					response.Json(ctx,err.ToR())
				}else{
					response.Json(ctx,response.InternalServerError)
				}
			}
		}
	}()
	ctx.Next()
}
