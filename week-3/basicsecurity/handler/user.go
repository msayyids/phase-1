package handler

import "database/sql"

type UserHandler struct {
	DB *sql.DB
}

func (u *UserHandler)