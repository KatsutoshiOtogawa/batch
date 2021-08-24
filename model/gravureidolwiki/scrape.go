package gravureidolwiki

// ユーザーに関する処理のエントリーポイントです。
import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// 指定のxpathに対して、url一覧を返す。
func FetchGravureIdorUriFromSite(xpath string) ([]string, error) {

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
	// page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(os.TempDir()),
	// , chromedp.NodeVisible
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		// navigate pornhub
		chromedp.Navigate(HomePageUrl),

		// 指定のあいうえおの行一覧に移動
		chromedp.Click(xpath, chromedp.BySearch),

		// 指定のあいうえおの行のものを取ってくる。
		chromedp.WaitReady("//*[@id=\"left\"]/ul/li"),
		chromedp.Nodes("//*[@id=\"left\"]/ul/li", &nodes),

		chromedp.Screenshot("/html", &imageBuf, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
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
	return nil, nil
}
