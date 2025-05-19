package models

import (
	"errors"

	"github.com/RevanthGovindan/event-booking/db"
	"github.com/RevanthGovindan/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email,password) VALUES(?,?)`
	statement, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id,password FROM users WHERE email = ?`
	var hashPassword string
	row := db.Db.QueryRow(query, u.Email)
	err := row.Scan(&u.ID, &hashPassword)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(u.Password, hashPassword) {
		return errors.New("invalid credentials")
	}

	return nil
}
