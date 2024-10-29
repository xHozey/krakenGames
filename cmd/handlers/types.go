package handler

import (
	"database/sql"
	"mime/multipart"
)

type Post struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Content string   `json:"content"`
	Screens []string `json:"screens"`
	System  SystemR  `json:"system"`
	Info    GameInfo `json:"info"`
}

type SystemR struct {
	OS        string `json:"os"`
	Processor string `json:"processor"`
	Memory    string `json:"memory"`
	Graphics  string `json:"graphics"`
	DirectX   string `json:"direct_x"`
	Storage   string `json:"storage"`
}

type GameInfo struct {
	Genre     string `json:"genre"`
	Developer string `json:"developer"`
	Platform  string `json:"platform"`
	GameSize  string `json:"game_size"`
	Released  string `json:"released"`
	Version   string `json:"version"`
}

type PostForm struct {
	Title   string                  `form:"title"`
	Image   *multipart.FileHeader   `form:"image"`
	Content string                  `form:"content"`
	Screens []*multipart.FileHeader `form:"screens"`
	System  SystemRForm             `form:"system"`
	Info    GameInfoForm            `form:"info"`
}

type SystemRForm struct {
	OS        string `form:"system[os]"`
	Processor string `form:"system[processor]"`
	Memory    string `form:"system[memory]"`
	Graphics  string `form:"system[graphics]"`
	DirectX   string `form:"system[direct_x]"`
	Storage   string `form:"system[storage]"`
}

type GameInfoForm struct {
	Genre        string `form:"info[genre]"`
	Developer    string `form:"info[developer]"`
	Platform     string `form:"info[platform]"`
	GameSize     string `form:"info[gameSize]"`
	ReleasedDate string `form:"info[released]"`
	Version      string `form:"info[version]"`
}

type DataBase struct {
	Db *sql.DB
}
