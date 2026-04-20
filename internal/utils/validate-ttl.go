package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ValidateTtl(ttlStr string) (time.Duration, error) {
	// default to 1 day
	if ttlStr == "" {
		return 1 * 24 * time.Hour, nil
	}

	if !strings.HasSuffix(ttlStr, "d") {
		return 0, errors.New("invalid ttl format - expected format is <number>d (e.g. 1d, 30d)")
	}

	daysStr := strings.TrimSuffix(ttlStr, "d")
	days, err := strconv.Atoi(daysStr)
	if err != nil {
		return 0, errors.New("invalid ttl - must be between 1d and 30d")
	}

	if days < 1 || days > 30 {
		return 0, errors.New("invalid ttl - must be between 1d and 30d")
	}

	return time.Duration(days) * 24 * time.Hour, nil
}
