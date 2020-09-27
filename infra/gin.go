package infra

import (
	"github.com/alancesar/account/middleware"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	return engine
}

func StartServer(engine *gin.Engine) {
	engine.Run()
}