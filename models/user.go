package models

import (
	"api-rest/db"
	"api-rest/utils"
	"errors"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	// Query para inserir dados e retornar o ID gerado
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	hashedPassword, error := utils.HashPassword(user.Password)
	if error != nil {
		return error
	}

	// Executa a query e captura o ID gerado
	err := db.DB.QueryRow(query, user.Email, hashedPassword).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u User) ValidateCredentials() error {
	// Query com placeholder do PostgreSQL
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		// Se não encontrar um resultado, ou ocorrer outro erro
		return errors.New("credentials invalid")
	}

	// Verifica a senha utilizando a função utilitária
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
