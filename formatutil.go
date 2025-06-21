// Package mediafileinfo
package mediafileinfo

import (
	"fmt"
	"math"
)

// FormatBytes converts an int64 value of bytes into a human-readable string using KB, MB, GB, or TB (1024 basis).
// The result is rounded to 2 decimals (e.g., 1536 -> 1.50 KB).
func FormatBytes(bytes int64) string {
	const (
		_          = iota
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
	)

	b := float64(bytes)
	round2 := func(val float64) float64 {
		return math.Round(val*100) / 100
	}

	switch {
	case b >= TB:
		return fmt.Sprintf("%.2f TB", round2(b/TB))
	case b >= GB:
		return fmt.Sprintf("%.2f GB", round2(b/GB))
	case b >= MB:
		return fmt.Sprintf("%.2f MB", round2(b/MB))
	case b >= KB:
		return fmt.Sprintf("%.2f KB", round2(b/KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

// FormatDurationMS converts time in milliseconds to a human-readable string.
// Hides hours if zero, hides minutes if both hours and minutes are zero.
func FormatDurationMS(ms uint64) string {
	hours := ms / 3600000
	ms = ms % 3600000
	minutes := ms / 60000
	ms = ms % 60000
	seconds := ms / 1000
	ms = ms % 1000

	switch {
	case hours > 0:
		return fmt.Sprintf("%d:%02d:%02d.%03d", hours, minutes, seconds, ms)
	case minutes > 0:
		return fmt.Sprintf("%d:%02d.%03d", minutes, seconds, ms)
	default:
		return fmt.Sprintf("%d.%03d", seconds, ms)
	}
}
