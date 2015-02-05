package gigasecond

import "time"

// TestVersion
const TestVersion = 1

// Birthday my birthday
var Birthday, _ = time.Parse("2006-01-02", "1987-06-01")

// AddGigasecond adds 1 gigasecond into input and returns the result
func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Second * 1e9)
}
