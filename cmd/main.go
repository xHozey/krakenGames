package main

import (
	handler "kraken/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../assets/templates/*")
	r.Static("/assets/design", "../assets/design")
	r.GET("/", handler.HomeHandler)
	r.Run()
}
