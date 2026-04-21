# Go Screenshot

A simple API to render high-quality website screenshots. <br />
Built using go and [chromedp](https://github.com/chromedp/chromedp)

## API

https://go-screenshot.onrender.com/ss?url=https://github.com/trending

#### Query Parameters

| Parameter | Required | Default   | Description                                       |
| --------- | -------- | --------- | ------------------------------------------------- |
| `url`     | Yes      | -         | The page URL to capture                           |
| `device`  | No       | `desktop` | Device preset - `desktop` or `mobile`             |
| `ttl`     | No       | `1d`      | Cache duration in days ranging from `1d` to `30d` |
