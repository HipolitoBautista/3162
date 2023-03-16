package models

import (
	"database/sql"
	"time"
)

type form struct {
	UserID                   int32
	user_id                  int32
	form_id                  int32
	form_status              string
	archive_status           bool
	affiant_full_name        string
	other_names              string
	name_change_status       string
	social_security_num      int32
	social_security_date     time.Time
	social_security_country  string
	passport_number          int32
	passport_date            time.Time
	passport_country         string
	dob                      time.Time
	place_of_birth           string
	nationality              string
	acquired_nationality     string
	spouse_name              string
	affiants_address         string
	residencial_phone_number int32
	residenceial_fax_num     int32
	residencial_email        string
	created_on               time.Time
}

// setup dependency injection
type FormModel struct {
	DB *sql.DB //connection pool
}
