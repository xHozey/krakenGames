package handler

import "database/sql"

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
	Genre        []string `json:"genre"`
	Developer    string   `json:"developer"`
	Platform     string   `json:"platform"`
	GameSize     string   `json:"game_size"`
	ReleasedDate string   `json:"released_date"`
	Version      string   `json:"version"`
}

type DataBase struct {
	Db *sql.DB
}
