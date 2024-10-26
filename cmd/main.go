package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../assets/templates/*")
	r.Static("/assets/design", "../assets/design")
	r.GET("/")
	r.Run()
}
