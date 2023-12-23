package handler

import (
	"context"
	"errors"
	"pairproject/entity"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(email, password string) (entity.Users, error) {
	user := entity.Users{}
	ctx := context.Background()
	query := `
		SELECT userId, name, age, email, password
		FROM Users WHERE Email = ?
	`

	rows, err := h.DB.QueryContext(ctx, query, email)
	if err != nil {
		return user, err
	}

	if rows.Next() {
		err := rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Email, &user.Password)
		if err != nil {
			return user, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return user, errors.New("email/password invalid")
		}
	}
	return user, nil
}
