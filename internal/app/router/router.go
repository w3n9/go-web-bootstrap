package router

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	routeV1(r.Group("/api/v1"))
}
