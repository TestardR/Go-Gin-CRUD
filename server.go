package main

import (
	"io"
	"os"

	"github.com/TestardR/Go-Gin-CRUD/controller"
	"github.com/TestardR/Go-Gin-CRUD/middlewares"
	"github.com/TestardR/Go-Gin-CRUD/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	// TODO: Use the file if already exists
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	// TODO: Add PUT and DELETE route
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save((ctx)))
	})

	server.Run(":8080")
}
