package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ResolveChromePath() string {
	for _, envVar := range []string{"CHROME_BIN", "GOOGLE_CHROME_BIN", "CHROMIUM_PATH", "CHROME_PATH"} {
		if path := strings.TrimSpace(os.Getenv(envVar)); path != "" && chromePathExists(path) {
			return path
		}
	}

	for _, candidate := range []string{
		".render/chrome/opt/google/chrome/google-chrome",
		"./.render/chrome/opt/google/chrome/google-chrome",
		"/opt/render/project/src/.render/chrome/opt/google/chrome/google-chrome",
		"/opt/render/project/.render/chrome/opt/google/chrome/google-chrome",
		"google-chrome",
		"chromium",
		"chromium-browser",
		"chrome",
		"/usr/bin/google-chrome",
		"/usr/bin/chromium",
		"/usr/bin/chromium-browser",
		"/snap/bin/chromium",
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		`C:\Program Files\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
	} {
		if resolvedPath := resolveBrowserCandidate(candidate); resolvedPath != "" {
			return resolvedPath
		}
	}

	return ""
}

func resolveBrowserCandidate(candidate string) string {
	if filepath.IsAbs(candidate) || strings.Contains(candidate, "/") || strings.Contains(candidate, `\`) {
		if chromePathExists(candidate) {
			return candidate
		}
		return ""
	}

	resolvedPath, err := exec.LookPath(candidate)
	if err != nil {
		return ""
	}

	return resolvedPath
}

func chromePathExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir()
}
