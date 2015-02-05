package clock

import "fmt"

// TestVersion
const TestVersion = 1

// Number of minutes in a day
const DayMinutes = 60 * 24

// Clock is a struct with hours and minutes
type Clock struct{ h, m int }

// New returns a new clock instance
func New(h, m int) Clock {
	t := ((h*60+m)%DayMinutes + DayMinutes) % DayMinutes
	return Clock{h: t / 60, m: t % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add adds a number of minutes and return a new Clock instance
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}
