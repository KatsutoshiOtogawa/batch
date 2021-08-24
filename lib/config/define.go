package config

import (
	"database/sql"
)

type Args struct {
	FuncName    string
	AlphabetFlg uint
	Db          *sql.DB
}
