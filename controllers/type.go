package controllers

import (
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
