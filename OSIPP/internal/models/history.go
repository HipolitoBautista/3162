package models

import (
	"database/sql"
	"time"
)

type Form struct {
	admin_id  int32
	comments  string
	edit_made time.Time
}

// setup dependency injection
type formModel struct {
	DB *sql.DB //connection pool
}
