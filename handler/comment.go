package handler

import (
	"net/http"

	"project.dev/api/db"

	"github.com/gin-gonic/gin"
)

type CommentItem struct {
	ID       *int   `form:"id" json:"id"`
	Sentence string `form:"sentence" json:"sentence"`
	Fullname string `form:"fullname" json:"fullname"`
	Job      string `form:"job" json:"job"`
	Post     int    `form:"post" json:"post"`
}

func CreateComment(c *gin.Context) {

	var form CommentItem
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error

	// Create comment user
	newComment := db.Comment{
		Sentence:  form.Sentence,
		Fullname:  form.Fullname,
		Job:       form.Job,
		PostingID: int(form.Post),
	}
	err = db.DBConn.Create(&newComment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfully created",
	})
}

func GetComment(c *gin.Context) {

	var listComment []db.Comment

	// Get comment's users
	err := db.DBConn.
		Where("posting_id = ?", c.Params.ByName("postingId")).
		Find(&listComment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": listComment})

}
