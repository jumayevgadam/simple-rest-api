package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/simple-rest-api/controller"
	"github.com/jumayevgadam/simple-rest-api/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
	routeURL                                   = "/videos"
)

func main() {
	app := gin.Default()

	app.GET(routeURL, func(ctx *gin.Context) {
		videos := VideoController.FindAll()
		ctx.JSON(200, videos)
	})

	app.POST(routeURL, func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	app.Run(":8080")
}
