package utils

import (
	"os"
	"runtime"
	"strings"
)

func isProduction() bool {
	return strings.EqualFold(strings.TrimSpace(os.Getenv("IS_PRODUCTION")), "true")
}

func chromePathExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func ResolveChromePath() string {
	if !isProduction() || runtime.GOOS != "linux" {
		return ""
	}

	if path := strings.TrimSpace(os.Getenv("CHROME_BIN")); path != "" && chromePathExists(path) {
		return path
	}

	if chromePathExists(RenderChromePath) {
		return RenderChromePath
	}

	return ""
}
