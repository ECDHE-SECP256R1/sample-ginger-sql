package models

import "github.com/kulichak/sql"

type Hello struct {
	sql.BaseModel

	FirstName string
	LastName  string
}
