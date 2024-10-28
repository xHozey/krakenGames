package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (db *DataBase) Post(g *gin.Context) {
	if g.Request.Method == http.MethodGet {
		g.HTML(http.StatusOK, "post.html", nil)
	} else if g.Request.Method == http.MethodPost {
		title := g.PostForm("title")
		content := g.PostForm("content")
		imageFile, err := g.FormFile("image")
		if err != nil {
			fmt.Println("Error retrieving image:", err)
			g.String(http.StatusBadRequest, "Error retrieving image")
			return
		}
		imgPath := uploadImage(imageFile, g)
		form, err := g.MultipartForm()
		if err != nil {
			fmt.Println("Error retrieving multipart form:", err)
			g.String(http.StatusBadRequest, "Error retrieving multipart form")
			return
		}
		screenFiles := form.File["screens"]

		db.insertData(title, content, imgPath, screenFiles, g)
	}
}

func (db *DataBase) Api(g *gin.Context) {
	posts, err := db.getData()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	g.IndentedJSON(http.StatusOK, posts)
}
