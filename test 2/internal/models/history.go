package models

import (
	"database/sql"
	"time"
)

type history struct {
	admin_id  int32
	comments  string
	edit_made time.Time
}

// setup dependency injection
type HistoryModel struct {
	DB *sql.DB //connection pool
}



