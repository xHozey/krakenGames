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
	r.Static("/static", "../assets/frontend/build/static")
	r.StaticFile("/", "../assets/frontend/build/index.html")
	r.Static("/uploads", "../assets/uploads")
	r.GET("/api", data.GetPosts)
	r.GET("/api/:id", data.GetPostById)
	r.GET("/post", data.Post)
	r.POST("/post", data.Post)
	r.Run()
}
