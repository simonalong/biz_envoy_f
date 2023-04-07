package main

import (
	"biz-f-service/biz_config"
	"biz-f-service/router"

	"github.com/isyscore/isc-gobase/server"
)

func main() {
	biz_config.InitConfig()

	router.Register()

	server.Run()
}
