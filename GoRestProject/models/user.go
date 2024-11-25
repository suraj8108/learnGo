package models

import (
	"database/sql"
	"errors"

	"github.com/suraj8108/learngo/Gorestproject/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save(db *sql.DB) error {
	query := "INSERT INTO usersEvents (email, password) VALUES ($1, $2)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateCredentials(db *sql.DB) error {
	query := `
		SELECT id, password from usersEvents where email = $1
	`

	row := db.QueryRow(query, u.Email)

	var retrievedPassowrd string

	err := row.Scan(&u.ID, &retrievedPassowrd)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(u.Password, retrievedPassowrd) {
		return errors.New("Incorrect Password")
	}
	return nil

}
