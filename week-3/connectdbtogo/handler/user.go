package handler

import (
	"connectdb/entity"
	"context"
	"database/sql"
)

type UserHandler struct {
	DB *sql.DB
}



func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) FindAll() ([]entity.User, error) {
	var users []entity.User
	ctx := context.Background()
	rows, err := h.DB.QueryContext(ctx, "SELECT id, name, age FROM Users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
