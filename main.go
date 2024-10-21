package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"project.dev/api/db"
	"project.dev/api/handler"
)

func init() {
	godotenv.Load()
	db.Init()
	db.Migrate(db.DBConn)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/create/post", handler.CreatePost)
	r.GET("/list/post", handler.GetPost)
	r.POST("/create/comment", handler.CreateComment)
	r.GET("/list/comment/:postingId", handler.GetComment)
	r.POST("/add/like", handler.AddLike)
	r.GET("/count/like/:postingId", handler.GetLike)

	r.Run()
}
