package main

import (
	"database/sql"
	"fmt"

	handler "kraken/cmd/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/Data.db")
	if err != nil {
		fmt.Println("error in opening db: ", err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("../assets/templates/*")
	r.Static("/design", "../assets/design")
	r.Static("/scripts", "../assets/scripts")

	r.GET("/games", handler.Api)
	r.GET("/post", handler.Post)
	r.POST("/post", handler.Post)
	r.GET("/", handler.HomeHandler)
	r.Run()
}
