package controller

import (
	"example/web-api/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
}