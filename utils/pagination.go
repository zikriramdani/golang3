package utils

import (
	"golang3/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	// Initializing default
	//	var mode string
	limit := 8
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
	}
}
