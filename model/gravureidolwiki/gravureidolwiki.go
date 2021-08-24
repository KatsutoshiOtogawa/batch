package gravureidolwiki

// ユーザーに関する処理のエントリーポイントです。
import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/KatsutoshiOtogawa/batch/lib/config"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

//　Mock
func Mock(args *config.Args) error {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("no-sandbox", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	start := time.Now()
	// navigate to a page, wait for an element, click

	var res string
	var nodes []*cdp.Node

	var imageBuf []byte
	// macの横幅、縦幅
	const (
		MacBookProWidth  = 2880
		MacBookProHeight = 1800
	)

	width, height := MacBookProWidth, MacBookProHeight

	// , chromedp.NodeVisible
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		// navigate pornhub
		chromedp.Navigate(HomePageUrl),

		// あ行
		chromedp.Click("//*[@id=\"sidemenu\"]/ul[1]/li[2]/a", chromedp.BySearch),

		// わ行
		// chromedp.Click("//*[@id=\"sidemenu\"]/ul[1]/li[11]/a", chromedp.BySearch),

		// あ行の最初
		// chromedp.Click("//*[@id=\"left\"]/ul/li[1]/a", chromedp.BySearch),

		chromedp.WaitReady("//*[@id=\"left\"]/ul/li"),
		chromedp.Nodes("//*[@id=\"left\"]/ul/li", &nodes),

		// chromedp.Click("//*[@id=\"signinSubmit\"]", chromedp.BySearch),

		// でバックログ
		// chromedp.CaptureScreenshot(&imageBuf),
		//*[@id="usernameModal"]
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.ScrollIntoView("footer"),
		// // chromedp.WaitVisible(`footer < div`),
		// chromedp.Text("h1", &res, chromedp.NodeVisible, chromedp.ByQuery),
		// chromedp.WaitVisible(`h1`, chromedp.ByQuery),
		chromedp.Screenshot("/html", &imageBuf, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	t := time.Now()
	if err := ioutil.WriteFile(fmt.Sprintf("./log/screenshot/%4d-%02d-%2d %02d:%02d:%02d %d.png", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()), imageBuf, 0644); err != nil {
		log.Fatal(err)
	}

	for _, li := range nodes {
		a := li.Children
		href := a[0].AttributeValue("href")
		fmt.Printf("node: %s | href = %s\n", li.LocalName, href)
	}

	fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
	return nil
}

//　あ行のグラビアアイドル
func あ行のグラビアアイドル(args *config.Args) ([]string, error) {

	uris, err := FetchGravureIdorUriFromSite("//*[@id=\"sidemenu\"]/ul[1]/li[2]/a")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return uris, nil
}

//　か行のグラビアアイドル
func か行のグラビアアイドル(args *config.Args) ([]string, error) {

	uris, err := FetchGravureIdorUriFromSite("//*[@id=\"sidemenu\"]/ul[1]/li[3]/a")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return uris, nil
}

//　さ行のグラビアアイドル
func さ行のグラビアアイドル(args *config.Args) ([]string, error) {

	uris, err := FetchGravureIdorUriFromSite("//*[@id=\"sidemenu\"]/ul[1]/li[4]/a")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return uris, nil
}

// グラビアアイドルのサイトをスクレイプする。
func FetchGravureIdorUri(args *config.Args) ([]string, error) {

	var gravureIdorUri []string

	// 引数より使う変数。
	alphabetFlg := (*args).AlphabetFlg

	if alphabetFlg&Theあcolumn == Theあcolumn {
		あ, err := あ行のグラビアアイドル(args)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		gravureIdorUri = append(gravureIdorUri, あ...)
	}

	if alphabetFlg&Theかcolumn == Theかcolumn {
		か, err := か行のグラビアアイドル(args)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		gravureIdorUri = append(gravureIdorUri, か...)
	}

	if alphabetFlg&Theさcolumn == Theさcolumn {
		さ, err := さ行のグラビアアイドル(args)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		gravureIdorUri = append(gravureIdorUri, さ...)
	}

	return gravureIdorUri, nil
}

// グラビアアイドルのUrlをもとにそのページの情報を取得。
// おそらくtableに入れれるレベルの構造体で返すのが作りとして正しい。
func FetchGravureIdolInfo(args *config.Args) (*ArticleInfo, error) {

	articleInfo := new(ArticleInfo)
	sss, err := FetchGravureIdolThumbnailFromSite("http://smarturl.it/uchidario-wiki")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	(*articleInfo).ThumbnailUrl = sss[0]

	sss, err = FetchGravureIdolProfileFromSite("http://smarturl.it/uchidario-wiki")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return articleInfo, nil
	// var gravureIdorUri []string

	// // 引数より使う変数。
	// alphabetFlg := (*args).AlphabetFlg

	// if alphabetFlg&Theあcolumn == Theあcolumn {
	// 	あ, err := あ行のグラビアアイドル(args)

	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return nil, err
	// 	}

	// 	gravureIdorUri = append(gravureIdorUri, あ...)
	// }

	// if alphabetFlg&Theかcolumn == Theかcolumn {
	// 	か, err := か行のグラビアアイドル(args)

	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return nil, err
	// 	}

	// 	gravureIdorUri = append(gravureIdorUri, か...)
	// }

	// if alphabetFlg&Theさcolumn == Theさcolumn {
	// 	さ, err := さ行のグラビアアイドル(args)

	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return nil, err
	// 	}

	// 	gravureIdorUri = append(gravureIdorUri, さ...)
	// }

	// return gravureIdorUri, nil
}
