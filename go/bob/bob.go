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
	remark = strings.TrimSpace(remark)
	switch {
	case strings.Contains(remark, " I "):
		return "Whoa, chill out!"
		// 匹配句子中的所有为大写的字母
	case strings.ContainsFunc(remark, func(rune) bool {
		if strings.IndexFunc(remark, unicode.IsLetter) == -1 {
			return false
		}
		return strings.ToUpper(remark) == remark
	}):
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
