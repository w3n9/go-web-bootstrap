package health

import (
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/pkg/common/response"
	"github.com/gin-gonic/gin"
)

type x struct{
	X string `json:"x"`
}

func Ping(ctx *gin.Context) {
	//_,err:=ioutil.ReadFile("")
	//response.AbortWithError(ctx,err)
	response.SuccessWithNoData(ctx)
}
