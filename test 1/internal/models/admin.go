package models

import (
	"context"
	"database/sql"
	"time"
)

type admin struct {
	id          int32
	email       string
	au_password int32
	created_at  time.Time
}

// setup dependency injection
type AdminModel struct {
	DB *sql.DB //connection pool
}

func (m *AdminModel) InsertAdmins() (int64, error) {
	var id int64
	email := "hipbautest@gmail.com"
	password := "testpass1234"
	statement := `
	INSERT INTO admin_users(email, au_password)
	VALUES($1, $2)
	RETURNING id
 `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, email, password).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err

}
