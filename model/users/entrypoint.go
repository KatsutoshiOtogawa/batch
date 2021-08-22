package users

// ユーザーに関する処理のエントリーポイントです。
import (
	"fmt"

	"github.com/KatsutoshiOtogawa/batch/lib/config"
)

// 各関数へのエントリーポイント
func Invoke(args *config.Args) {

	switch (*args).FuncName {

	case "Mock":
		Mock(args)

	default:
		fmt.Println((*args).FuncName, "は存在しない関数です。")
	}
}
