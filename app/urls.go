package app

import (
	"github.com/gin-gonic/gin"
	"github.com/soshika/sample-search/controllers/SE"
)

func urlPatterns() {

	authorized := router.Group("/api/v1", gin.BasicAuth(gin.Accounts{
		"api-key": "L1^hGk8z!jT6oD2p",
	}))

	authorized.POST("/search", SE.Search)
	authorized.POST("/index", SE.IndexExcel)
}
