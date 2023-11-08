package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {

	SetTimeZone()
	urlPatterns()
	router.Use(CORSMiddleware())

	router.Run(":9099")
}
