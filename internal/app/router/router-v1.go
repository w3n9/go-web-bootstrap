package router

import (
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/controller/health"
	"github.com/gin-gonic/gin"
)

func routeV1(r *gin.RouterGroup){
	r.GET("/ping",health.Ping)
}