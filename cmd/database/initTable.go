package database

import (
	"database/sql"
	"fmt"
)

func InitDB(db *sql.DB) error {
	// Create Post table
	createPostTable := `CREATE TABLE IF NOT EXISTS Post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		image TEXT,
		content TEXT
	);`

	// Create SystemR table
	createSystemRTable := `CREATE TABLE IF NOT EXISTS SystemR (
		post_id INTEGER,
		os TEXT,
		processor TEXT,
		memory TEXT,
		graphics TEXT,
		direct_x TEXT,
		storage TEXT,
		FOREIGN KEY (post_id) REFERENCES Post(id) ON DELETE CASCADE
	);`

	// Create GameInfo table
	createGameInfoTable := `CREATE TABLE IF NOT EXISTS GameInfo (
		post_id INTEGER,
		genre TEXT,
		developer TEXT,
		platform TEXT,
		game_size TEXT,
		released_date TEXT,
		version TEXT,
		FOREIGN KEY (post_id) REFERENCES Post(id) ON DELETE CASCADE
	);`

	// Create Screens table for storing multiple screenshots per Post
	createScreensTable := `CREATE TABLE IF NOT EXISTS Screens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		screen TEXT,
		FOREIGN KEY (post_id) REFERENCES Post(id) ON DELETE CASCADE
	);`

	// Execute table creation statements
	if _, err := db.Exec(createPostTable); err != nil {
		return fmt.Errorf("failed to create Post table: %v", err)
	}
	if _, err := db.Exec(createSystemRTable); err != nil {
		return fmt.Errorf("failed to create SystemR table: %v", err)
	}
	if _, err := db.Exec(createGameInfoTable); err != nil {
		return fmt.Errorf("failed to create GameInfo table: %v", err)
	}
	if _, err := db.Exec(createScreensTable); err != nil {
		return fmt.Errorf("failed to create Screens table: %v", err)
	}

	return nil
}
