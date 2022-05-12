package models

import "database/sql"

type User struct {
	Id       sql.NullInt32  `json:"id"`
	Username sql.NullString `json:"username"`
	Password sql.NullString `json:"password"`
}
