package http

import (
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/config"
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/middleware"
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/router"
	"github.com/gin-gonic/gin"
)

func Run(){
	r := gin.Default()
	r.Use(middleware.BodyCollector)
	r.Use(middleware.GlobalErrorHandler)
	router.Route(r)
	_ = r.Run(config.Application.Server.Host + ":" + config.Application.Server.Port)
}
