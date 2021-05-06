package health

import (
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/pkg/common/response"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	response.Json(ctx, response.Success)
}
