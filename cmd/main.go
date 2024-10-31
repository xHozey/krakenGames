package main

import (
	"database/sql"
	"fmt"

	database "kraken/cmd/database"
	handler "kraken/cmd/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database/kraken.db")
	data := handler.DataBase{Db: db}
	database.InitDB(db)
	if err != nil {
		fmt.Println("error in opening db: ", err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("../assets/templates/*")
	r.Static("/scripts", "../assets/scripts")
	r.Static("/uploads", "../assets/uploads")
	r.Static("/design", "../assets/design")
	r.GET("/", data.HomeHandler)
	r.GET("/api", data.GetPosts)
	r.GET("/api/:id", data.GetPostById)
	r.GET("/post", data.Post)
	r.POST("/post", data.Post)
	r.GET("/:id", data.GameInfo)
	r.Run()
}
