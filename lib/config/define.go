package config

import (
	"database/sql"
)

type Args struct {
	FuncName string
	Db       *sql.DB
}
