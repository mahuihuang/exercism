// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
// type Kind

type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral 等边
	Iso             // isosceles 等腰
	Sca             // scalene 不等边
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind
	if a != 0 && a+b >= c && b+c >= a && a+c >= b {
		switch {
		case a == b && b == c:
			k = Equ
		case (a == b && b != c) || (a == c && a != b) || (b == c && a != b):
			k = Iso
		default:
			k = Sca
		}
	} else {
		k = NaT
	}

	return k
}
