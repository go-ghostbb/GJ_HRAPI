package htmltopdf

import (
	"context"
	"fmt"
	"ghostbb.io/gb/frame/g"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"net"
	"time"
)

func Create(url string, sel string, timeout time.Duration) ([]byte, error) {
	// 如果需要debug解開底下註釋，並將ctx傳入，即可看到視窗
	//options := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	//ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)

	taskCtx, cancel := chromedp.NewContext(
		//ctx,
		context.Background(),
		// option
		chromedp.WithLogf(Printf),
	)
	defer cancel()

	if timeout != 0 {
		taskCtx, cancel = context.WithTimeout(taskCtx, timeout)
		defer cancel()
	}

	var pdfBuffer []byte
	if err := chromedp.Run(taskCtx, pdfGrabber(url, sel, &pdfBuffer)); err != nil {
		return nil, err
	}

	return pdfBuffer, nil
}

func pdfGrabber(url string, sel string, res *[]byte) chromedp.Tasks {

	start := time.Now()
	return chromedp.Tasks{
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		// chromedp.Text(`h1`, &res, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 1920 * 1080
			// .WithPaperHeight(12.8).WithPaperWidth(7.2)
			buf, _, err := page.PrintToPDF().WithPaperHeight(12.8).WithPaperWidth(7.2).WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			//fmt.Printf("h1 contains: '%s'\n", res)
			fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())
			return nil
		}),
	}
}

// 檢查是否有啟用chromedp/headless-shell
func checkChromePort(host string, port int) bool {
	addr := net.JoinHostPort(host, gbconv.String(port))
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func Printf(format string, v ...any) {
	g.Log().Infof(gbctx.New(), format, v...)
}
