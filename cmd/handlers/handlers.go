package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db *DataBase) Post(g *gin.Context) {
	if g.Request.Method == http.MethodGet {
		g.File("../assets/frontend/build/index.html")
		return
	} else if g.Request.Method == http.MethodPost {
		form, err := g.MultipartForm()
		if err != nil {
			fmt.Println(err)
		}
		title := form.Value["title"][0]
		image := form.File["image"][0]
		imgPath := uploadImage(image, g)
		content := form.Value["content"][0]
		genre := form.Value["info[genre]"][0]
		developer := form.Value["info[developer]"][0]
		platform := form.Value["info[platform]"][0]
		gameSize := form.Value["info[gameSize]"][0]
		released := form.Value["info[released]"][0]
		version := form.Value["info[version]"][0]
		screens := form.File["screens"]
		os := form.Value["system[os]"][0]
		processor := form.Value["system[processor]"][0]
		directX := form.Value["system[direct_x]"][0]
		graphics := form.Value["system[graphics]"][0]
		memory := form.Value["system[memory]"][0]
		storage := form.Value["system[storage]"][0]
		db.insertData(title, content, imgPath, genre, developer, platform, gameSize, released, version, os, processor, directX, graphics, memory, storage, screens, g)
	}
}

func (db *DataBase) GetPosts(g *gin.Context) {
	posts := db.getData()
	g.IndentedJSON(http.StatusOK, posts)
}

func (db *DataBase) GetPostById(g *gin.Context) {
	id := g.Param("id")
	post := db.getDataByID(id)
	g.IndentedJSON(http.StatusOK, post)
}
