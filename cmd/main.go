package main

import (
	"fmt"
	"jamm3e3333/file-upload/cmd/app/config"
	"jamm3e3333/file-upload/cmd/internal"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gd := gin.Default()

	storageCfg := config.NewGcsCfg()
	appCfg := config.NewAppCfg()
	internal.Register(gd, storageCfg)

	err := gd.Run(fmt.Sprintf(":%d", appCfg.Port))
	if err != nil {
		log.Fatalf("failed to start the http server: %s", err.Error())
	}
}
