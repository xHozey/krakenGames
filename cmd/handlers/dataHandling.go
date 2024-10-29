package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (db *DataBase) insertData(title, content, imgPath, genre, developer, platform, gameSize, released, version, os, processor, directX, graphics, memory, storage string, screenFiles []*multipart.FileHeader, g *gin.Context) {
	stm, err := db.Db.Prepare("INSERT INTO Post (title, image, content) VALUES (?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stm.Exec(title, imgPath, content)
	if err != nil {
		fmt.Println(err)
	}
	postID, _ := res.LastInsertId()
	for _, screen := range screenFiles {
		screenPath := uploadImage(screen, g)
		stm, err = db.Db.Prepare("INSERT INTO Screens (post_id, screen) VALUES (?,?)")
		if err != nil {
			fmt.Println(err)
		}
		stm.Exec(postID, screenPath)
	}
	stm, err = db.Db.Prepare("INSERT INTO GameInfo (post_id, genre, developer, platform, game_size, released, version) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	stm.Exec(postID, genre, developer, platform, gameSize, released, version)

	stm, err = db.Db.Prepare("INSERT INTO SystemR (post_id, os, processor, memory, graphics, direct_x, storage) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	stm.Exec(postID, os, processor, memory, graphics, directX, storage)
}

func (db *DataBase) getData() []Post {
	var games []Post
	rows, err := db.Db.Query("SELECT id, title, image, content FROM Post")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for rows.Next() {
		var game Post
		rows.Scan(&game.Id, &game.Title, &game.Image, &game.Content)
		screenRow, _ := db.Db.Query("SELECT screen FROM Screens WHERE post_id = ?", game.Id)
		for screenRow.Next() {
			var screen string
			screenRow.Scan(&screen)
			game.Screens = append(game.Screens, screen)
		}
		db.Db.QueryRow("SELECT os, processor, memory, graphics, direct_x, storage FROM SystemR WHERE post_id = ?", game.Id).Scan(&game.System.OS, &game.System.Processor, &game.System.Memory, &game.System.Graphics, &game.System.DirectX, &game.System.Storage)
		db.Db.QueryRow("SELECT genre, developer, platform, game_size, released, version FROM GameInfo WHERE post_id = ?", game.Id).Scan(&game.Info.Genre, &game.Info.Developer, &game.Info.Platform, &game.Info.GameSize, &game.Info.Released, &game.Info.Version)
		games = append(games, game)
	}
	return games
}

func (db *DataBase) getDataByID(id string) Post {
	var game Post
	db.Db.QueryRow("SELECT id, title, image, content FROM Post WHERE id = ?", id).Scan(&game.Id, &game.Title, &game.Image, &game.Content)
	db.Db.QueryRow("SELECT os, processor, memory, graphics, direct_x, storage FROM SystemR WHERE post_id = ?", id).Scan(&game.System.OS, &game.System.Processor, &game.System.Memory, &game.System.Graphics, &game.System.DirectX, &game.System.Storage)
	db.Db.QueryRow("SELECT genre, developer, platform, game_size, released, version FROM GameInfo WHERE post_id = ?", id).Scan(&game.Info.Genre, &game.Info.Developer, &game.Info.Platform, &game.Info.GameSize, &game.Info.Released, &game.Info.Version)
	return game
}

func uploadImage(img *multipart.FileHeader, g *gin.Context) string {
	imagePath := filepath.Join("../assets/uploads", img.Filename)

	if err := g.SaveUploadedFile(img, imagePath); err != nil {
		fmt.Println("Error saving image:", err)
		g.String(http.StatusInternalServerError, "Error saving image")
		return ""
	}
	return imagePath
}
