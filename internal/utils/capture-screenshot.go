package utils

import (
	"context"
	"errors"
	"go-screenshot/internal/models"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func CaptureScreenshot(urlStr, deviceStr string) ([]byte, error) {
	userAgent := DesktopUserAgent
	isMobile := false

	if deviceStr == models.DeviceMobile {
		userAgent = MobileUserAgent
		isMobile = true
	}

	width, height := GetWidthHeight(deviceStr)

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.UserAgent(userAgent),
			chromedp.Flag("headless", true),
			chromedp.Flag("hide-scrollbars", true),
			chromedp.Flag("mute-audio", true),
		)...,
	)
	defer cancelAllocator()

	ctx, cancelCtx := chromedp.NewContext(allocatorCtx)
	defer cancelCtx()

	ctx, cancelTimeout := context.WithTimeout(ctx, Timeout)
	defer cancelTimeout()

	var screenshotBytes []byte
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride(userAgent),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1, isMobile),
		chromedp.Navigate(urlStr),
		chromedp.Sleep(SleepDuration),
		chromedp.CaptureScreenshot(&screenshotBytes),
	)
	if err != nil {
		return nil, errors.New("failed to capture screenshot: " + err.Error())
	}

	return screenshotBytes, nil
}
