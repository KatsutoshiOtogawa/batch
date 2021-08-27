package users

// ユーザーに関する処理のエントリーポイントです。
import (
	"database/sql"
	"log"

	"github.com/KatsutoshiOtogawa/batch/lib/config"
)

//　Mock
func Mock(args *config.Args) error {
	db := (*args).Db

	stmt, err := db.Prepare(`
	select 'Hello World'
	`)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

// ログインが許可されているユーザーかどうか
func PermittedLoginUser(username string, password string, db *sql.DB) (bool, error) {

	//
	stmt, err := db.Prepare(`
	select ?,?
	`)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	_, err = stmt.Exec(
		username,
		password,
	)
	if err != nil {
		log.Println(err.Error())

	}

	return true, nil
}
