package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status checks the status of the service
func (h HealthController) Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
