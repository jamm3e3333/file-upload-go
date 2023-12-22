package main

import (
	"jamm3e3333/file-upload/application/handler"
	"jamm3e3333/file-upload/ui/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gd := gin.Default()

	c := controller.NewFileUpload(handler.Handle)
	c.Register(gd)
	err := gd.Run(":8889")

	if err != nil {
		log.Fatal(err)
	}
}
