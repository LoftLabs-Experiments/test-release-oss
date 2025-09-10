package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GenerateID generates a unique identifier
func GenerateID() string {
	return uuid.New().String()
}

// FormatTimestamp formats a timestamp for display
func FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// SanitizeName sanitizes a name by removing special characters
func SanitizeName(name string) string {
	// Replace spaces and special characters with hyphens
	sanitized := strings.ReplaceAll(name, " ", "-")
	sanitized = strings.ToLower(sanitized)
	return sanitized
}

// BuildInfo returns build information
func BuildInfo(version, buildTime string) string {
	return fmt.Sprintf("Version: %s, Built: %s", version, buildTime)
}
