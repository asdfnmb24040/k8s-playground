package server

import (
	"log"
	"go-project/controller"
	"go-project/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinServer() {
	r := setRouter()
	err := r.Run(utils.GetEnv("GIN_SERVER_IP"))

	if err != nil {
		log.Fatalf("upload server run fail")
		return
	}
}

func setRouter() *gin.Engine {
	r := gin.Default()

	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders: []string{"authorize", "Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method",
			"Access-Control-Request-Headers"},
	}

	r.Use(cors.New(corsConf))
	r.Use(middleware)
	r.GET("/GinEcho", controller.GinEcho)
	r.GET("/GinRedisInit", controller.GinRedisInit)
	r.GET("/GinRedisEcho", controller.GinRedisEcho)

	return r
}

func middleware(ctx *gin.Context) {
	// check token
	ctx.Next()
}
