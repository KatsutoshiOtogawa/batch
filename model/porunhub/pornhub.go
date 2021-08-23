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

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	start := time.Now()
	// navigate to a page, wait for an element, click
	// https://jp.pornhub.com/view_video.php?viewkey=ph5ddf2ce189a7d
	var res string

	var imageBuf []byte

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),

		// navigate pornhub
		chromedp.Navigate("https://jp.pornhub.com"),
		// porunhub_login
		chromedp.Click("//*[@id=\"headerLoginLink\"]", chromedp.BySearch),
		chromedp.SendKeys("//*[@id=\"usernameModal\"]", pornhub_usernmae, chromedp.BySearch),
		chromedp.SendKeys("//*[@id=\"passwordModal\"]", pornhub_password, chromedp.BySearch),
		chromedp.Click("//*[@id=\"signinSubmit\"]", chromedp.BySearch),

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
	if err := ioutil.WriteFile(fmt.Sprintf("%4d%2d%2d %2d:%2d%2d %d.png", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()), imageBuf, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
	return nil
}
