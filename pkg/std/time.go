package std

import "time"

//CreateTime creates a timeImpl instance.
func CreateTime() *timeImpl {
	return &timeImpl{}
}

//timeImpl is a wrapper of the time standard package.
type timeImpl struct{}

//Now returns the current time.
func (t *timeImpl) Now() time.Time {
	return time.Now()
}
