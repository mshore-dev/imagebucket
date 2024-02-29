package database

import "time"

type Album struct {
	ID          int
	Shortcode   string
	Name        string
	Description string
	CreatedAt   time.Time
	CreatedBy   int
	Public      bool
}

func CreateAlbum(shortcode, name, description string, createdby int, public bool) error {

	return nil
}

func DeleteAlbum(id int) error {

	return nil
}

func GetAlbumsByUser(id, limit, offset int) ([]Album, error) {

	return []Album{}, nil
}

func GetAlbumByShortcode(code string) (Album, error) {

	return Album{}, nil
}

func GetAlbumFileCount(id int) (int, error) {

	return 0, nil
}

// ? all these
// func GetAlbumContents()

// func AddFileToAlbum()

// func RemoveFileFromAlbum()
