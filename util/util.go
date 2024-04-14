package util

import (
	"time"
	"unicode/utf8"
)

func TruncateText(text string, maxChars int) string {
	if utf8.RuneCountInString(text) <= maxChars {
		return text
	}
	return text[:maxChars-3] + "..."
}

func FormatDate(date time.Time) string {

	return date.Format("2006-01-02")

}
