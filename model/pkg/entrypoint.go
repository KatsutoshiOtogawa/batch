package pkg

// ユーザーに関する処理のエントリーポイントです。
import (
	"fmt"

	"github.com/KatsutoshiOtogawa/batch/lib/config"
	"github.com/KatsutoshiOtogawa/batch/model/gravureidolwiki"
	"github.com/KatsutoshiOtogawa/batch/model/pornhub"
	"github.com/KatsutoshiOtogawa/batch/model/users"
)

// 各パッケージへのinvoke
func Invoke(pkgName string, args *config.Args) {
	switch pkgName {
	case "users":
		users.Invoke(args)

	case "pornhub":
		pornhub.Invoke(args)

	case "gravureidolwiki":
		gravureidolwiki.Invoke(args)
	default:
		fmt.Println(pkgName, "は存在しないパッケージです。")
	}

}
