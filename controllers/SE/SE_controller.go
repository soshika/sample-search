package SE

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/soshika/sample-search/domains/SE"
	"github.com/soshika/sample-search/logger"
	"github.com/soshika/sample-search/services"
	"net/http"
	"path/filepath"
)

var (
	dist = "/home/soshika/projects/sample-search/uploads/"
)

func IndexExcel(c *gin.Context) {
	logger.Info("Enter to to IndexExcel controller successfully")

	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension
	fullPath := dist + newFileName

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	excel := SE.Excel{
		FileName: fullPath,
		Index:    "excel",
	}

	data, serviceErr := services.SEService.IndexExcel(&excel)
	if serviceErr != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)

	logger.Info("Close from IndexExcel controller successfully")
}

func Search(c *gin.Context) {
	logger.Info("Enter to Search controller successfully")

	searchREQ := SE.SearchEngineReq{}
	if err := c.ShouldBindJSON(&searchREQ); err != nil {
		logger.Error("error when trying to bind json", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	data, serviceErr := services.SEService.Search(&searchREQ)
	if serviceErr != nil {

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)

	logger.Info("Close from Search controller successfully")
}
