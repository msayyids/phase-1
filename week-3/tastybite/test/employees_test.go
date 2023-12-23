package test

import (
	"database/sql"
	"tastybite/config"
	"tastybite/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db          *sql.DB
	UserHandler handler.UserHandler
)

func TestMain(m *testing.M) {

	db = config.ConnectDb()

	UserHandler = handler.UserHandler{DB: db}

	m.Run()

	db.Exec("truncate table employees")
}

func TestAddEmp(t *testing.T) {
	t.Run("berhasil menambah employee", func(t *testing.T) {
		err := UserHandler.AddEmployees("ahmad", "anu", "chef")
		assert.NoError(t, err)
	})

	t.Run("gagal menambah employee", func(t *testing.T) {
		err := UserHandler.AddEmployees("", "", "")
		assert.Error(t, err)
	})
}
