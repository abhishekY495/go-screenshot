package utils

import (
	"context"
	"errors"
	"go-screenshot/internal/models"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func CaptureScreenshot(request models.ScreenshotRequest) ([]byte, error) {
	userAgent := DesktopUserAgent
	isMobile := false

	if request.Device == models.DeviceMobile {
		userAgent = MobileUserAgent
		isMobile = true
	}

	width, height := GetWidthHeight(request.Device)

	allocatorOptions := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(userAgent),
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("mute-audio", true),
	)

	if chromePath := ResolveChromePath(); chromePath != "" {
		allocatorOptions = append(allocatorOptions, chromedp.ExecPath(chromePath))
	}

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(
		context.Background(),
		allocatorOptions...,
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
		chromedp.Navigate(request.Url),
		chromedp.Sleep(SleepDuration),
		chromedp.CaptureScreenshot(&screenshotBytes),
	)
	if err != nil {
		return nil, errors.New("failed to capture screenshot: " + err.Error())
	}

	return screenshotBytes, nil
}
