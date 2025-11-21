package logs

import (
	"fmt"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	application := "default"
	for _, char := range log {
		switch fmt.Sprintf("%U", char) {
		case "U+2757":
			application = "recommendation"
		case "U+1F50D":
			application = "search"
		case "U+2600":
			application = "weather"
		default:
		}
		if application != "default" {
			break
		}
	}
	return application
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, fmt.Sprintf("%c", oldRune), fmt.Sprintf("%c", newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := 0
	for range log {
		length++
	}
	return length <= limit
}
