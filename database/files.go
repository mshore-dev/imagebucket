package database

import "time"

type File struct {
	ID               int
	Filename         string
	OriginalFilename string
	CreatedAt        int64
	CreatedBy        int
	Hash             string
}

func GetFileByID(id int) (File, error) {

	var f File

	row := db.QueryRow("SELECT * FROM files WHERE id = ?", id)

	err := row.Scan(&f.ID, &f.Filename, &f.OriginalFilename, &f.CreatedAt, &f.CreatedBy, &f.Hash)
	if err != nil {
		return File{}, nil
	}

	return f, nil
}

func GetFileByName(filename string) (File, error) {

	var f File

	row := db.QueryRow("SELECT * FROM files WHERE filename = ?", filename)

	err := row.Scan(&f.ID, &f.Filename, &f.OriginalFilename, &f.CreatedAt, &f.CreatedBy, &f.Hash)
	if err != nil {
		return File{}, nil
	}

	return f, nil
}

// maybe?
// func GetFileByShortcode()

func GetFileByHash(hash string) (File, error) {

	var f File

	row := db.QueryRow("SELECT * FROM files WHERE hash = ?", hash)

	err := row.Scan(&f.ID, &f.Filename, &f.OriginalFilename, &f.CreatedAt, &f.CreatedBy, &f.Hash)
	if err != nil {
		return File{}, nil
	}

	return f, nil
}

func GetFilesByUser(userid, limit, offset int) ([]File, error) {

	// TODO: verify limit and offset somehow?

	rows, err := db.Query("SELECT * FROM files WHERE createdby = ? LIMIT ? OFFSET ?", userid, limit, offset)
	if err != nil {
		// TODO: check if error is "no rows"?
		return []File{}, err
	}

	var files []File
	var f File

	for rows.Next() {
		err := rows.Scan(&f.ID, &f.Filename, &f.OriginalFilename, &f.CreatedAt, &f.CreatedBy, &f.Hash)
		if err != nil {
			return []File{}, err
		}

		files = append(files, f)
	}

	return files, nil
}

func CreateFile(filename, originalFilename, hash string, createdby int) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO files (filename, originalfilename, createdat, createdby, hash) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(filename, originalFilename, time.Now().Unix(), createdby, hash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// TODO: deleting files
// func DeleteFile(id int) error {

// 	return nil
// }
