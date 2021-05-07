package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func BodyCollector(ctx *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
	ctx.Set("RequestBody", string(bodyBytes))
	_ = ctx.Request.Body.Close()
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
}
