package main

import (
	"io"
	"net/http"
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

	router := gin.New()
	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	// TODO: Add PUT and DELETE route
	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	router.GET("/videos/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		video := videoController.FindOne(id)
		ctx.JSON(200, gin.H{"message": video})
	})

	router.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Input is valid"})
		}
	})

	router.Run(":8080")
}
