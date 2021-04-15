package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"opb/models"
	"strconv"
)

type UsersController struct {
	db *gorm.DB
}

var user models.User
var users []models.User

func NewUsersController(db *gorm.DB) ControllerInterface {
	var users ControllerInterface
	users = &UsersController{db: db}
	return users
}

func (this *UsersController) Index(c *gin.Context) {
	this.db.Order("created_at desc").Find(&users)
	c.JSON(http.StatusOK, users)
}

func (this *UsersController) Show(c *gin.Context) {
	this.db.First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)
}

func (this *UsersController) Create(c *gin.Context) {
	c.Bind(&user)
	this.db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (this *UsersController) Update(c *gin.Context) {
	c.Bind(&user)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user.ID = uint(id)
	this.db.Save(&user)
	c.JSON(http.StatusCreated, user)
}

func (this *UsersController) Delete(c *gin.Context) {
	this.db.First(&user, c.Param("id"))
	this.db.Delete(&user)
	c.JSON(http.StatusNoContent, gin.H{})
}
