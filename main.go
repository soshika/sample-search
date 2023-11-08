package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soshika/sample-search/app"
)

func main() {
	// TODO: should change to release mode!
	gin.SetMode(gin.DebugMode)
	app.StartApplication()
}
