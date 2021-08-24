package gravureidolwiki

// ユーザーに関する処理のエントリーポイントです。
import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// ２階層
// 指定のxpathに対して、url一覧を返す。
func FetchGravureIdolThumbnailFromSite(navigateUrl string) ([]string, error) {

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
	var nodes []*cdp.Node

	// var imageBuf []byte
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

		chromedp.Navigate(navigateUrl),

		chromedp.WaitReady("//*[@id=\"left\"]/a/img", chromedp.BySearch),
		chromedp.Nodes("//*[@id=\"left\"]/a/img", &nodes, chromedp.BySearch),
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// サムネイルなので画像が一つしか無いので、決め打ち
	img := nodes[0]
	src := img.AttributeValue("src")
	fmt.Printf("node: %s | src = %s\n", img.LocalName, src)

	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
	return []string{src, ""}, nil
}

// ２階層
// プロファイル
func FetchGravureIdolProfileFromSite(navigateUrl string) ([]string, error) {

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
	// start := time.Now()
	// navigate to a page, wait for an element, click
	// var nodes []*cdp.Node
	var innerHtml string
	// var imageBuf []byte
	// macの横幅、縦幅
	const (
		MacBookProWidth  = 2880
		MacBookProHeight = 1800
	)

	width, height := MacBookProWidth, MacBookProHeight
	// <div id="infobox">朝比奈祐未（あさひな ゆみ) プロフィール<br>
	// 愛称	ひな
	// ひなけつ<br>
	// 生年月日	1991年12月9日<br>

	// 出身地	日本・神奈川県横浜市<br>
	// 血液型	B型<br>
	// 身長 / 体重	158 cm / ― kg<br>
	// スリーサイズ	89 - 59 - 93 cm<br>
	// カップサイズ	G<br>
	// </div>
	// 	内田理央 プロフィール<br>
	// 別生年月日      1991年9月27日<br>
	// 出身地   日本・東京都<br>
	// 血液型  O型<br>
	// 公称サイズ（2010年時点）<br>
	// 身長 / 体重     166 cm / ― kg<br>
	// スリーサイズ    80 - 58 - 82 cm<br>

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),

		chromedp.Navigate(navigateUrl),

		chromedp.WaitReady("//*[@id=\"infobox\"]", chromedp.BySearch),
		// chromedp.Nodes("//*[@id=\"infobox\"]", &nodes, chromedp.BySearch),
		chromedp.InnerHTML("//*[@id=\"infobox\"]", &innerHtml, chromedp.BySearch),
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// サムネイルなので画像が一つしか無いので、決め打ち
	// div := nodes[0]
	// src := div.ContentDocument.ContentDocument
	// fmt.Printf("node: %s | src = %s\n", img.LocalName, src)

	type Profile struct {
		BirthPlace string
		BirthDate  string
	}

	profile := new(Profile)

	profiles := strings.Split(innerHtml, "<br>")
	for _, p := range profiles {

		if strings.Contains(p, "生年月日") {
			p = strings.Replace(
				strings.Replace(p, " ", "", 0),
				"生年月日",
				"",
				0,
			)

			// (*profile).BirthDate

		} else if strings.Contains(p, "出身地") {
			(*profile).BirthPlace = strings.Replace(
				strings.Replace(p, " ", "", 0),
				"出身地",
				"",
				0,
			)
		}
	}
	// fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
	return []string{innerHtml, ""}, nil
}
