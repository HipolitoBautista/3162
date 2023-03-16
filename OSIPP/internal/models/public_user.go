package models

import (
	"database/sql"
)

type PublicUser struct {
	id          int32
	email       string
	pu_password int32
}

// setup dependency injection
type publicUserModel struct {
	DB *sql.DB //connection pool
}
