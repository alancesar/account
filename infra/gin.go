package infra

import (
	"github.com/gin-gonic/gin"
)

func CreateServer(middlewares ...gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(middlewares...)
	return engine
}

func StartServer(engine *gin.Engine) {
	_ = engine.Run()
}
