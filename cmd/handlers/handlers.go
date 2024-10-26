package handler

import (
	"fmt"
	"net/http"
	"os"

	"kraken/cmd/types"

	"github.com/gin-gonic/gin"
)

func HomeHandler(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", nil)
}

func Post(g *gin.Context) {
	// if g.Request.Method == http.MethodGet {
	// 	g.HTML(http.StatusOK, "post.html", nil)
	// } else if g.Request.Method == http.MethodPost {
	// 	title := g.PostForm("title")
	// 	content := g.PostForm("content")
	// 	image, err := g.FormFile("image")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(0)
	// 	}
	// 	form, err := g.MultipartForm()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(0)
	// 	}
	// 	screenFiles := form.File["screens"]

	// }
}

func Api(g *gin.Context) {
	posts := []types.Post{
		{
			Id:      1,
			Title:   "The Legend of Zelda: Breath of the Wild",
			Image:   "https://example.com/zelda-cover.jpg",
			Content: "Explore a vast open world and save the kingdom of Hyrule.",
			Screens: []string{
				"https://example.com/zelda-screen1.jpg",
				"https://example.com/zelda-screen2.jpg",
				"https://example.com/zelda-screen3.jpg",
			},
			System: types.SystemR{
				OS:        "Nintendo Switch OS",
				Processor: "Custom NVIDIA Tegra processor",
				Memory:    "4GB LPDDR4",
				Graphics:  "NVIDIA Maxwell-based GPU",
				DirectX:   "N/A",
				Storage:   "32GB Internal",
			},
			Info: types.GameInfo{
				Genre:        []string{"Action", "Adventure"},
				Developer:    "Nintendo EPD",
				Platform:     "Nintendo Switch",
				GameSize:     "14.4GB",
				ReleasedDate: "2017-03-03",
				Version:      "1.6.0",
			},
		},
		{
			Id:      2,
			Title:   "Cyberpunk 2077",
			Image:   "https://example.com/cyberpunk-cover.jpg",
			Content: "An open-world, action-adventure story set in Night City, a megalopolis obsessed with power and body modification.",
			Screens: []string{
				"https://example.com/cyberpunk-screen1.jpg",
				"https://example.com/cyberpunk-screen2.jpg",
				"https://example.com/cyberpunk-screen3.jpg",
			},
			System: types.SystemR{
				OS:        "Windows 10",
				Processor: "Intel Core i7-4790 or AMD Ryzen 3 3200G",
				Memory:    "12GB",
				Graphics:  "NVIDIA GTX 1060 or AMD Radeon R9 Fury",
				DirectX:   "12",
				Storage:   "70GB SSD",
			},
			Info: types.GameInfo{
				Genre:        []string{"Action", "RPG"},
				Developer:    "CD Projekt Red",
				Platform:     "PC",
				GameSize:     "60GB",
				ReleasedDate: "2020-12-10",
				Version:      "1.63",
			},
		},
		{
			Id:      3,
			Title:   "God of War",
			Image:   "https://example.com/godofwar-cover.jpg",
			Content: "Journey with Kratos and his son Atreus through the Norse wilderness, facing gods and monsters alike.",
			Screens: []string{
				"https://example.com/godofwar-screen1.jpg",
				"https://example.com/godofwar-screen2.jpg",
				"https://example.com/godofwar-screen3.jpg",
			},
			System: types.SystemR{
				OS:        "Windows 10",
				Processor: "Intel Core i5-6600k or AMD Ryzen 5 2400 G",
				Memory:    "8GB",
				Graphics:  "NVIDIA GTX 1060 or AMD Radeon RX 570",
				DirectX:   "11",
				Storage:   "70GB SSD",
			},
			Info: types.GameInfo{
				Genre:        []string{"Action", "Adventure"},
				Developer:    "Santa Monica Studio",
				Platform:     "PC",
				GameSize:     "64GB",
				ReleasedDate: "2022-01-14",
				Version:      "1.0.12",
			},
		},
		{
			Id:      4,
			Title:   "Red Dead Redemption 2",
			Image:   "https://example.com/rdr2-cover.jpg",
			Content: "Experience the epic story of Arthur Morgan, a lawman-turned-outlaw, in Rockstar’s sprawling Wild West.",
			Screens: []string{
				"https://example.com/rdr2-screen1.jpg",
				"https://example.com/rdr2-screen2.jpg",
				"https://example.com/rdr2-screen3.jpg",
			},
			System: types.SystemR{
				OS:        "Windows 10",
				Processor: "Intel Core i7-4770K or AMD Ryzen 5 1500X",
				Memory:    "12GB",
				Graphics:  "NVIDIA GeForce GTX 1060 6GB or AMD Radeon RX 480 4GB",
				DirectX:   "12",
				Storage:   "150GB",
			},
			Info: types.GameInfo{
				Genre:        []string{"Action", "Adventure"},
				Developer:    "Rockstar Games",
				Platform:     "PC",
				GameSize:     "112GB",
				ReleasedDate: "2018-10-26",
				Version:      "1436.31",
			},
		},
	}
	g.IndentedJSON(http.StatusOK, posts)
}
