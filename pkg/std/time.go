package std

import "time"

//CreateTimeWrapper creates a TimeWrapper instance.
func CreateTimeWrapper() *TimeWrapper {
	return &TimeWrapper{}
}

//TimeWrapper is a wrapper of the time standard package.
type TimeWrapper struct{}

//Now returns the current time.
func (t *TimeWrapper) Now() time.Time {
	return time.Now()
}

//Sleep waits for a certain amount of time.
func (t *TimeWrapper) Sleep(d time.Duration) {
	time.Sleep(d)
}
