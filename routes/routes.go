package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"opb/controllers"
)

func Init(db *gorm.DB) *gin.Engine {
	var users controllers.ControllerInterface
	users = controllers.NewUsersController(db)

	router := gin.Default()
	router.GET("/users", users.Index)
	router.GET("/users/:id", users.Show)
	router.POST("/users", users.Create)
	router.PUT("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)

	return router
}
