package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mshore-dev/imagebucket/config"
)

var db *sql.DB

func OpenDB() {
	var err error

	db, err = sql.Open("sqlite3", config.Config.Database)
	if err != nil {
		log.Fatalf("failed to open database: %v\n", err)
	}

	// is there any benifit to having this in a seperate function?
	// this is just how I've always done things.
	createDB()
}

func createDB() {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS "users" (
		"ID"	INTEGER,
		"Username"	TEXT,
		"Password"	TEXT,
		PRIMARY KEY("ID" AUTOINCREMENT)
	);

	CREATE TABLE IF NOT EXISTS "files" (
		"ID"	INTEGER,
		"FileName"	REAL,
		"OriginalFileName"	TEXT,
		"CreatedAt"	INTEGER,
		"CreatedBy"	INTEGER,
		"Hash"	TEXT,
		FOREIGN KEY("CreatedBy") REFERENCES "users"("ID"),
		PRIMARY KEY("ID" AUTOINCREMENT)
	);

	CREATE TABLE IF NOT EXISTS "albums" (
		"ID"	INTEGER,
		"ShortCode"	TEXT,
		"Name"	TEXT,
		"Description"	TEXT,
		"CreatedAt"	INTEGER,
		"CreatedBy"	INTEGER,
		"Public"	INTEGER,
		PRIMARY KEY("ID" AUTOINCREMENT)
	);

	CREATE TABLE IF NOT EXISTS "albums_to_files" (
		"AlbumID"	INTEGER,
		"FileID"	INTEGER,
		FOREIGN KEY("FileID") REFERENCES "files"("ID"),
		FOREIGN KEY("AlbumID") REFERENCES "albums"("ID")
	);
	`)
	if err != nil {
		log.Fatalf("failed to run create query: %v\n", err)
	}

}
