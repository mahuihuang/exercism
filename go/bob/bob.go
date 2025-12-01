// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	trimRemark := strings.TrimSpace(remark)
	switch {
	case strings.Contains(trimRemark, " I "):
		return "Whoa, chill out!"
	case strings.ContainsFunc(trimRemark, func(rune) bool {
		flag := false
		for _, char := range trimRemark {
			if unicode.IsLetter(char) {
				if unicode.IsUpper(char) {
					flag = true
				} else {
					return false
				}
			}
		}
		return flag
	}):
		if strings.HasSuffix(trimRemark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case strings.HasSuffix(trimRemark, "?"):
		return "Sure."
	case trimRemark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
