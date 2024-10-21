package handler

import (
	"net/http"

	"gorm.io/gorm"
	"project.dev/api/db"

	"github.com/gin-gonic/gin"
)

type LikeItem struct {
	ID   *int `form:"id" json:"id"`
	Post int  `form:"post" json:"post"`
}

func AddLike(c *gin.Context) {

	var form LikeItem
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error

	// Add like by user
	newLike := db.Like{
		PostingID: int(form.Post),
	}
	err = db.DBConn.Create(&newLike).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Like has been added",
	})
}

type LikeData struct {
	TotalLike int64 `json:"totalLike"`
}

func GetLike(c *gin.Context) {

	var countLike LikeData

	// Counting like on post
	sql := db.DBConn.
		Model(&db.Like{}).
		Session(&gorm.Session{})

	sql.Where("posting_id = ?", c.Params.ByName("postingId")).Count(&countLike.TotalLike)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": countLike})

}
