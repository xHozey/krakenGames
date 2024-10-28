package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (db *DataBase) insertData(title, content, imgPath string, screenFiles []*multipart.FileHeader, g *gin.Context) {
	stm, err := db.Db.Prepare("INSERT INTO Post (title, content, image) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stm.Close()

	res, err := stm.Exec(title, content, imgPath)
	if err != nil {
		fmt.Println("Error executing insert statement:", err)
		return
	}
	postID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving last insert ID:", err)
		return
	}
	for _, screen := range screenFiles {
		screenPath := uploadImage(screen, g)
		stm, err := db.Db.Prepare("INSERT INTO Screens (post_id, screen) VALUES (?,?)")
		if err != nil {
			fmt.Println("Error inserting screen shots:", err)
			g.String(http.StatusBadRequest, "Error inserting screen shots")
			return
		}
		stm.Exec(postID, screenPath)
	}
}

func (db *DataBase) getData() ([]Post, error) {
	stm, err := db.Db.Prepare("SELECT id, title, content, image FROM Post")
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Image)
		//stm,err := db.Db.Prepare("SELECT screen FROM Screens WHERE post_id = ?")
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return posts, nil
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
