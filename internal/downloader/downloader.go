package downloader

import (
	"context"

	"github.com/chromedp/chromedp"
)

func DownloadPage(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML("html", &html),
	)
	if err != nil {
		return "", err
	}

	return html, nil
}
