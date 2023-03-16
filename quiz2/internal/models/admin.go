package models

import (
	"database/sql"
	"time"
)

type admin struct {
  id int32
  email string
  au_password int32
  created_at time.Time
}

// setup dependency injection
type AdminModel struct {
	DB *sql.DB //connection pool
}
