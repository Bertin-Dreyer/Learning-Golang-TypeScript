package clock

import "fmt"

// Define the Clock type here.
// Clock represents a time of day without a date.
type Clock struct {
    minutes int
}

func New(h, m int) Clock {
	total := (h * 60) + m
    total = ((total % 1440) + 1440) % 1440

    return Clock{minutes:total}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes + m)
}

func (c Clock) Subtract(m int) Clock {
    return New(0, c.minutes - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
