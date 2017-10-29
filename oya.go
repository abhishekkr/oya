package main

import (
	"fmt"

	oyaController "github.com/abhishekkr/oya/oyaController"

	golenv "github.com/abhishekkr/gol/golenv"
	gin "github.com/gin-gonic/gin"
)

var (
	/*
		HTTPAt specifies server's listen-at config, can be overridden by env var OYA_HTTP. Defaults to '':9000'.
	*/
	HTTPAt = golenv.OverrideIfEnv("OYA_HTTP", ":9000")
)

func main() {
	GinUp(HTTPAt)
	fmt.Println("bye .")
}

/*
ginCors to set required HTTP configs.
*/
func ginCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}

/*
ginHandleErrors to manage issues at server side.
*/
func ginHandleErrors(ctx *gin.Context) {
	ctx.Next()
	errorToPrint := ctx.Errors.ByType(gin.ErrorTypePublic).Last()
	if errorToPrint != nil {
		ctx.JSON(500, gin.H{
			"status":  500,
			"message": errorToPrint.Error(),
		})
	}
}

/*
GinUp maps all routing logic and starts server.
*/
func GinUp(listenAt string) {
	kube := oyaController.Kubernetes{}

	router := gin.Default()
	router.Use(ginHandleErrors)
	router.Use(ginCors())

	alpha := router.Group("/alpha")
	{
		alpha.GET("/kube/:type", kube.Create)
	}

	router.Run(listenAt)
}
