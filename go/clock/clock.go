package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	h += m / 60
	m %= 60
	if m < 0 {
		m += 60
		h--
	}

	h %= 24
	if h < 0 {
		h += 24
	}
	return Clock{
		h: h,
		m: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
