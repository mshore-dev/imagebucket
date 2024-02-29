package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func CreateUser(username, password string) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(id int) (User, error) {

	var u User

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	err := row.Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func CheckUserLogin(username, password string) (bool, error) {

	var u User

	row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)

	err := row.Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// we don't need to return this error to the caller
			return false, nil
		}

		return false, err
	}

	return false, nil
}

func UpdateUser(id int, user User) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE users SET username = ?, password = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Password, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashed), nil

}
