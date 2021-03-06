package pornhub

// ユーザーに関する処理のエントリーポイントです。
import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/KatsutoshiOtogawa/batch/lib/config"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

//　Mock
func Mock(args *config.Args) error {

	pornhub_usernmae := os.Getenv("PORNHUB_USERNAME")

	pornhub_password := os.Getenv("PORNHUB_PASSWORD")

	if pornhub_usernmae == "" || pornhub_password == "" {
		err := errors.New("You don't set PORNHUB_USERNAME or PORNHUB_PASSWORD.")

		log.Println(err.Error())
		return err
	}

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
	// https://jp.pornhub.com/view_video.php?viewkey=ph5ddf2ce189a7d
	var res string

	var imageBuf []byte

	// macの横幅、縦幅
	const (
		MacBookProWidth  = 2880
		MacBookProHeight = 1800
	)

	width, height := MacBookProWidth, MacBookProHeight

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		// navigate pornhub
		chromedp.Navigate(HomePageUrl),
		// porunhub_login
		chromedp.Click("//*[@id=\"headerLoginLink\"]", chromedp.BySearch),
		chromedp.SendKeys("//*[@id=\"usernameModal\"]", pornhub_usernmae, chromedp.NodeVisible, chromedp.BySearch),
		chromedp.SendKeys("//*[@id=\"passwordModal\"]", pornhub_password, chromedp.NodeVisible, chromedp.BySearch),
		chromedp.Click("//*[@id=\"signinSubmit\"]", chromedp.NodeVisible, chromedp.BySearch),

		// でバックログ
		chromedp.CaptureScreenshot(&imageBuf),
		//*[@id="usernameModal"]
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.ScrollIntoView("footer"),
		// // chromedp.WaitVisible(`footer < div`),
		// chromedp.Text("h1", &res, chromedp.NodeVisible, chromedp.ByQuery),

		// chromedp.WaitVisible(`h1`, chromedp.ByQuery),
		// chromedp.Screenshot(`h1`, &imageBuf, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	t := time.Now()
	if err := ioutil.WriteFile(fmt.Sprintf("./log/screenshot/%4d-%02d-%2d %02d:%02d:%02d %d.png", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()), imageBuf, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
	return nil
}

// https://idolwiki.web.fc2.com/sagyou.html
