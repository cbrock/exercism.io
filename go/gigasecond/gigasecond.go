import (
	"math"
	"time"
)

const gigasecond := math.Pow(10, 9)

// AddGigasecond adds one gigasecond to the given time object
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond) * time.Second)
}
