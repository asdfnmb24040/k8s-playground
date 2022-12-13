package controller

import (
	"go-project/utils"
	"go-project/utils/redisHelper"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GinEcho(ctx *gin.Context) {
	ctx.String(http.StatusOK, "gin echo call")
}

func GinRedisInit(ctx *gin.Context) {
	// init redis
	success := redisHelper.InitRedis()
	if !success {
		ctx.String(http.StatusOK, "InitRedis: fail")
		return
	} else {
		ctx.String(http.StatusOK, "InitRedis: success")
	}
}

func GinRedisEcho(ctx *gin.Context) {
	key := "test-" + uuid.New().String()
	val := uuid.New().String()
	redisHelper.SetRedis(key, val, 1)
	res := redisHelper.GetRedis(key)

	utils.PrintObj(ctx.ClientIP(), "gin echo ClientIP")

	ctx.String(http.StatusOK, "gin echo call, res ="+res)
}
