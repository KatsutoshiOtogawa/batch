package users

// ユーザーに関する処理のエントリーポイントです。
import (
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
