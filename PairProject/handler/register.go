package handler

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(name, email, password string, age int) error {

	ctx := context.Background()
	query := `
		INSERT INTO users (name, email, password,age)
		VALUES (?,?,?,?)
	`

	byteHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = h.DB.ExecContext(ctx, query, name, email, string(byteHashed), age)

	if err != nil {
		return err
	}

	return nil
}
