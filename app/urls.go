package app

import "github.com/gin-gonic/gin"

func urlPatterns() {

	authorized := router.Group("/api/v1", gin.BasicAuth(gin.Accounts{
		"api-key": "L1^hGk8z!jT6oD2p",
	}))

	authorized.POST("/search")
}
