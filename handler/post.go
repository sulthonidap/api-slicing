package handler

import (
	"net/http"

	"project.dev/api/db"

	"github.com/gin-gonic/gin"
)

type PostItem struct {
	ID      *int   `form:"id" json:"id"`
	Mind    string `form:"mind" json:"mind"`
	Picture string `form:"picture" json:"picture"`
}

func CreatePost(c *gin.Context) {

	var form PostItem
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error

	// Create post user
	newPost := db.Posting{
		Mind:    form.Mind,
		Picture: form.Picture,
	}
	err = db.DBConn.Create(&newPost).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post successfully created",
	})
}

func GetPost(c *gin.Context) {

	var listPost []db.Posting

	// Get post by users
	err := db.DBConn.
		Find(&listPost).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": listPost})

}
