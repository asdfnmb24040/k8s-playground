package main

import (
	"go-project/server"
	"go-project/utils"
	// "go-project/utils/redisHelper"
	// "fmt"
)

func main() {
	// load .env
	success := utils.InitEnv()
	if !success {
		return
	}

	server.InitGinServer()
}
